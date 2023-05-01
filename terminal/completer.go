package terminal

import (
	"context"
	"os"
	"path"
	"strings"

	"github.com/ansurfen/cushion/go-prompt"
	"github.com/ansurfen/cushion/utils"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"

	hi "hulo/sdk/go"
)

// Completer manage completion rules
type Completer struct {
	sysTbls  []string
	userTbls []string
	schemas  map[string]*Schema
	schema   string
	term     *HuloTerminal
}

func NewCompleter() *Completer {
	completer := &Completer{
		schemas: make(map[string]*Schema),
		schema:  "default",
	}
	defaultSchema := &Schema{
		user:      true,
		hisotry:   prompt.NewHistory(),
		highlight: make([]prompt.HighlightRule, 0),
		modes: []prompt.CompletionMode{
			{Name: "Syntax mode", Attr: prompt.ModeAttrNoIcon},
			{Name: "Router mode", Attr: prompt.ModeAttrNoDescription},
		},
		modeHandles: []prompt.Completer{
			completer.defaultSyntaxRule,
			completer.defaultRouterRule,
		},
	}
	completer.schemas["default"] = defaultSchema
	return completer
}

func (c *Completer) MarshalJSON() ([]byte, error) {
	return []byte(utils.JsonStr(utils.NewJsonObject(map[string]utils.JsonValue{
		"schema": utils.NewJsonString(c.schema),
	}))), nil
}

func (c *Completer) SetSchema(schema string) {
	c.schema = schema
}

func (c *Completer) Schema() string {
	return c.schema
}

// CurSchema returns current schema pointer
func (c *Completer) CurSchema() *Schema {
	if schema, ok := c.schemas[c.schema]; ok {
		return schema
	}
	return NewSchema()
}

func (c *Completer) SetHighlight(schema string, rules []prompt.HighlightRule) {
	if s, ok := c.schemas[schema]; ok {
		s.highlight = rules
	}
}

func (c *Completer) getTable(g bool) []string {
	if g {
		return append(c.sysTbls, c.userTbls...)
	}
	return c.sysTbls
}

// defaultSyntaxRule to offer user suggest of command
func (c *Completer) defaultSyntaxRule(doc prompt.Document) []prompt.Suggest {
	schema := c.CurSchema()
	in := doc.TextBeforeCursor()
	ins := splitPipe(in)
	if len(ins) > 1 {
		in = ins[len(ins)-1]
	}
	var s []prompt.Suggest
	in = utils.TrimMultiSpace(strings.TrimLeft(in, " "))
	subcmd := strings.Split(in, " ")
	args := &lua.LTable{}
	tips := &lua.LTable{}
	for i := 1; i <= len(c.getTable(schema.user)); i++ {
		tips.Insert(i, c.term.VM.GetGlobalVar(c.getTable(schema.user)[i-1]))
	}
	for i := 1; i <= len(subcmd); i++ {
		args.Insert(i, lua.LString(subcmd[i-1]))
	}
	c.term.VM.FastCallByParam("LoadRule", []lua.LValue{tips, args})
	c.term.VM.GetGlobalVar("Candidates").(*lua.LTable).ForEach(func(l1, l2 lua.LValue) {
		suggest := l2.(*lua.LTable)
		tmp := prompt.Suggest{
			Icon:        "",
			Text:        "",
			Description: "",
			Comment:     false,
		}
		if v := suggest.RawGetString("text"); v.Type() != lua.LTNil {
			tmp.Text = v.String()
		}
		if v := suggest.RawGetString("icon"); v.Type() != lua.LTNil {
			tmp.Icon = v.String()
		}
		if v := suggest.RawGetString("desc"); v.Type() != lua.LTNil {
			tmp.Description = v.String()
		}
		if v := suggest.RawGetString("comment"); v.Type() != lua.LTNil && v.String() == "true" {
			tmp.Comment = true
		}
		s = append(s, tmp)
	})
	c.term.VM.SetGlobalVar("Candidates", &lua.LTable{})
	return prompt.FilterHasPrefix(s, doc.GetWordBeforeCursor(), true)
}

// defaultRouterRule to offer user suggest of filepath
func (c *Completer) defaultRouterRule(doc prompt.Document) []prompt.Suggest {
	in := doc.TextBeforeCursor()
	ins := splitPipe(in)
	if len(ins) >= 1 {
		in = strings.TrimLeft(ins[len(ins)-1], " ")
	}
	if sf := strings.Split(in, " "); len(sf) > 0 {
		in = sf[len(sf)-1]
	}
	s := []prompt.Suggest{}
	lastindex := strings.LastIndex(in, "/")
	if lastindex == -1 || in[len(in)-1] == ' ' {
		in = "."
	} else {
		in = in[:lastindex]
	}
	files, _ := os.ReadDir(in)
	for _, file := range files {
		if in == "." {
			s = append(s, prompt.Suggest{Text: file.Name()})
		} else {
			s = append(s, prompt.Suggest{Text: path.Join(in+"/", file.Name())})
		}
	}
	return prompt.FilterHasPrefix(s, doc.GetWordBeforeCursor(), true)
}

// defaultRule returns callback which could offer suggest according to current shecma.
// If schema isn't exist or error occurs, it'll return empty suggest array.
func (c *Completer) defaultRule(doc prompt.Document) []prompt.Suggest {
	mode := doc.GetMode()
	schema := c.CurSchema()
	if mode >= 0 && mode < len(schema.modeHandles) {
		return schema.modeHandles[mode](doc)
	}
	return []prompt.Suggest{}
}

// newSchema create and store schema when it isn't exist
func (c *Completer) newSchema(name string, schema *Schema) {
	if _, ok := c.schemas[name]; !ok {
		c.schemas[name] = schema
	} else {
		c.term.IO.WriteStderr(c.term.Logger.Error("schema is exist already"))
	}
}

func (c *Completer) NewRemoteSchema(name string, service string) {
	cli := c.term.Daemon.GetService(name)
	c.term.Completer.newSchema(name, &Schema{
		user:    false,
		hisotry: prompt.NewHistory(),
		modes: []prompt.CompletionMode{
			{Name: "Syntax mode"},
		},
		modeHandles: []prompt.Completer{
			func(doc prompt.Document) []prompt.Suggest {
				ctx := context.Background()
				r, err := cli.Completion(ctx, &hi.CompletionRequest{Str: doc.TextBeforeCursor()})
				if err != nil {
					panic(err)
				}
				ret := []prompt.Suggest{}
				for _, rd := range r.Suggests {
					ret = append(ret, prompt.Suggest{Comment: rd.Comment, Text: rd.Text, Description: rd.Description})
				}
				return prompt.FilterHasPrefix(ret, doc.GetWordBeforeCursor(), true)
			},
		},
	})
}

func (c *Completer) NewSchema(opt *lua.LTable) {
	type Opt struct {
		Name  string
		User  bool
		Fn    *lua.LFunction
		Rules []prompt.HighlightRule
	}
	var o Opt
	if err := gluamapper.Map(opt, &o); err != nil {
		c.term.IO.WriteStderr(err.Error())
	}
	c.newSchema(o.Name, &Schema{
		user:      o.User,
		hisotry:   prompt.NewHistory(),
		highlight: o.Rules,
		modes: []prompt.CompletionMode{
			{Name: "Syntax mode"},
		},
		modeHandles: []prompt.Completer{
			func(d prompt.Document) []prompt.Suggest {
				c.term.VM.EvalFunc(o.Fn, []lua.LValue{
					luar.New(c.term.VM.Interp(), c.term),
				})
				s := []prompt.Suggest{}
				for _, v := range c.getTable(o.User) {
					s = append(s, prompt.Suggest{Text: v})
				}
				return s
			},
		},
	})
}

// Schema manage and converge different Completer, just like namesapce
type Schema struct {
	// user determine whether command of user level enable
	user        bool
	hisotry     *prompt.History
	highlight   []prompt.HighlightRule
	modes       []prompt.CompletionMode
	modeHandles []prompt.Completer
}

func NewSchema() *Schema {
	return &Schema{
		user:      true,
		hisotry:   prompt.NewHistory(),
		highlight: []prompt.HighlightRule{},
		modes:     []prompt.CompletionMode{},
	}
}

// RegisterMode to add mode to schema
func (s *Schema) RegisterMode(mode prompt.CompletionMode, handle prompt.Completer) {
	s.modes = append(s.modes, mode)
	s.modeHandles = append(s.modeHandles, handle)
}

// UnregisterMode to remove specify mode
func (s *Schema) UnregisterMode(mode string) {
	for i := 0; i < len(s.modes); i++ {
		if s.modes[i].Name == mode {
			s.modes = append(s.modes[:i], s.modes[i+1:]...)
			s.modeHandles = append(s.modeHandles[:i], s.modeHandles[i+1:]...)
		}
	}
}

func (s *Schema) SetHighlight(rules []prompt.HighlightRule) {
	s.highlight = rules
}
