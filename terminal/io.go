package terminal

import (
	"fmt"
	"strings"

	"github.com/ansurfen/cushion/go-prompt"
)

// HuloIO manage input/ouput on console
type HuloIO struct {
	stdout    string
	stderr    string
	cmds      chan string
	Prefix    string
	Cursor    byte
	pw        *PathWalk
	term      *HuloTerminal
	interrupt bool
}

func NewHuloIO() *HuloIO {
	return &HuloIO{
		stdout:    "",
		Prefix:    "/",
		Cursor:    '>',
		cmds:      make(chan string, 10),
		pw:        NewPathWalk(true),
		interrupt: false,
	}
}

func (io *HuloIO) WriteStderr(data string) {
	io.stderr += data
}

func (io *HuloIO) WriteStdout(data string) {
	io.stdout += data
}

func (io *HuloIO) printStdout() {
	fmt.Print(io.stdout)
	io.stdout = ""
}

func (io *HuloIO) printStderr() {
	fmt.Print(io.stderr)
	io.stderr = ""
}

func (io *HuloIO) Read() string {
	if !io.interrupt {
		if io.term.Completer.schema == "default" {
			io.WriteStdout(fmt.Sprintf("\n%s%c", io.pw.Pwd(), io.Cursor))
		} else {
			io.WriteStdout(fmt.Sprintf("\n%s%c", io.Prefix, io.Cursor))
		}
	}
	io.printStderr()
	io.printStdout()
	if io.interrupt {
		return <-io.cmds
	}
	schema := io.term.Completer.CurSchema()

	io.cmds <- prompt.Input("", io.term.Completer.defaultRule,
		prompt.OptionTitle("Hulo"),
		prompt.OptionRegisterMode(schema.modes),
		prompt.OptionHistoryInstance(schema.hisotry),
		prompt.OptionColors(io.term.Patiner.render),
		prompt.OptionHighlight(schema.highlight, func(s string) string {
			return strings.ToLower(s)
		}))
	return <-io.cmds
}

func splitPipe(script string) []string {
	subscript := []string{}
	ban := false
	str := ""
	for _, ch := range script {
		if ch == '"' || ch == '\'' {
			ban = !ban
		}
		if ch != '|' || ban {
			str += string(ch)
		} else {
			subscript = append(subscript, str)
			str = ""
		}
	}
	if len(str) > 0 {
		subscript = append(subscript, str)
	}
	return subscript
}
