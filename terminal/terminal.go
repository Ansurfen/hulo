package terminal

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"os/signal"

	"github.com/ansurfen/cushion/runtime"
	"github.com/ansurfen/cushion/utils"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// HuloTerminal manage and schedule whole REPL (read-eval-print loop)
type HuloTerminal struct {
	Env       *HuloEnv               `json:"env"`
	IO        *HuloIO                `json:"io"`
	Completer *Completer             `json:"completer"`
	Daemon    *HuloDaemon            `json:"-"`
	Reg       *HuloRegistry          `json:"-"`
	VM        runtime.VirtualMachine `json:"-"`
	Logger    *HuloLogger            `json:"-"`
	Patiner   *HuloPatiner           `json:"-"`
	EventBus  *HuloEventBus          `json:"-"`
}

func NewHulo() *HuloTerminal {
	term := &HuloTerminal{
		Env:       NewHuloEnv(),
		IO:        NewHuloIO(),
		Reg:       NewHuloRegistry(),
		Patiner:   NewHuloPainter(),
		Completer: NewCompleter(),
		Logger:    NewHuloLogger(),
		VM:        runtime.NewVirtualMachine().Default(),
		EventBus:  NewHuloEventBus(),
	}
	term.inject()

	// set hook to load plugin on runtime
	term.VM.SafeSetGlobalFn(runtime.LuaFuncs{
		// Boot is used to load and execute loader located in {UserHome}/.hulo/loader
		"Boot": huloBoot(term),
		// Export to import completion data
		// you can find it in {UserHome}/.hulo/complete
		"Export": huloExport(term),
	})

	// register HuloScheduler for lazy restart
	HuloScheduler = map[uint8]func(*HuloTerminal) error{
		HSMountPlugin: hsMountPlugin,
		HSLoadTheme:   hsLoadTheme,
		HSInitDaemon:  hsInitDaemon,
	}
	return term
}

// inject is used to inject terminal
func (term *HuloTerminal) inject() {
	term.IO.term = term
	term.Reg.term = term
	term.Completer.term = term
	term.Logger.term = term
	term.EventBus.term = term
}

// toString is used to serialize HuloTermnial and returns `{}` if serialization fails.
func (term *HuloTerminal) toString() string {
	str, err := json.Marshal(term)
	if err != nil {
		return "{}"
	}
	return string(str)
}

// Initialize to init total operator of hulo
func (term *HuloTerminal) Initialize() {
	for _, e := range HuloScheduler {
		e(term)
	}
}

func (term *HuloTerminal) RegisterCmd(name string, fun any) {
	term.Reg.RegisterCmd(name, CallbackCommand{
		fn: func(args CommandArgv) {
			fun.(func(...interface{}) []interface{})(args.rawArgv)
		},
	})
	flag := true
	for _, tbl := range term.Completer.userTbls {
		if tbl == name {
			flag = false
			break
		}
	}
	if flag {
		term.VM.Eval(runtime.LuaAssignLR(runtime.LuaIdent(name), runtime.LuaMap(runtime.Luamap{
			"name": runtime.LuaString(name),
		})))
		term.Completer.userTbls = append(term.Completer.userTbls, name)
	}
}

func lvalue2Lfunction(v lua.LValue) *lua.LFunction {
	switch vv := v.(type) {
	case *lua.LFunction:
		return vv
	}
	return &lua.LFunction{}
}

func huloBoot(term *HuloTerminal) func(lvm *lua.LState) int {
	return func(lvm *lua.LState) int {
		name := lvm.CheckString(1)
		lvm.CheckTable(2).ForEach(func(funName, fun lua.LValue) {
			switch funName.String() {
			case "init":
				if err := lvm.CallByParam(lua.P{
					Fn: lvalue2Lfunction(fun),
				}, luar.New(lvm, term)); err != nil {
					panic(err)
				}
			case "mount":
				term.Reg.RegisterCmd("mysql", LuaCommand{
					fn:     lvalue2Lfunction(fun),
					loader: name,
				})
			case "destory":
			}
		})
		return 0
	}
}

func huloExport(term *HuloTerminal) func(lvm *lua.LState) int {
	return func(lvm *lua.LState) int {
		term.VM.SetGlobalVar(lvm.CheckString(1), lvm.CheckTable(2))
		flag := true
		name := lvm.CheckString(1)
		for _, tbl := range term.Completer.userTbls {
			if tbl == name {
				flag = false
			}
		}
		if flag {
			if name == "hulo" {
				term.Completer.sysTbls = append(term.Completer.sysTbls, "hulo")
			} else {
				term.Completer.userTbls = append(term.Completer.userTbls, name)
			}
		}
		return 0
	}
}

// Read returns user input from IO
func (term *HuloTerminal) Read() string {
	return term.IO.Read()
}

// Exec to execute command to be parsed and return signal
func (term *HuloTerminal) Exec(argv CommandArgv) Signal {
	cmd := term.Reg.Query(argv.cmd)
	out, err := cmd.Exec(term, argv)
	if err != nil {
		switch err.Error() {
		case Restart:
			return Signal{
				Op: SJMP,
			}
		case SetInterrupt:
			term.IO.interrupt = false
			return Signal{
				Op: STI,
			}
		case ClearInterrupt:
			term.IO.interrupt = true
			// term.IO.cmds <- "echo"
			// term.IO.cmds <- "echo"
			// term.IO.cmds <- "open"
			return Signal{
				Op: CLI,
			}
		case NotFoundCmd:
			if term.Env.Global {
				if err := term.exec(argv.raw); err != nil {
					term.IO.WriteStderr(err.Error())
				}
			}
		default:
			term.IO.WriteStderr(err.Error())
		}
	}
	term.IO.WriteStdout(out)
	return Signal{
		Op: NOP,
	}
}

// exec automatically fit in os enviroment to execute command.
// windows 10+ -> powershell, others -> cmd;
// linux, darwin -> /bin/bash
// It look like cushion.utils.Exec, but it redirect stdin, stdout, stderr
func (term *HuloTerminal) exec(raw string) error {
	var cmd *exec.Cmd
	switch utils.CurPlatform.OS {
	case "windows":
		switch utils.CurPlatform.Ver {
		case "10", "11":
			cmd = exec.Command("powershell", raw)
		default:
			cmd = exec.Command("cmd", []string{"/C", raw}...)
		}
	case "linux", "darwin":
		cmd = exec.Command("/bin/bash", []string{"/C", raw}...)
	default:
		return errors.New(UusupportPlatform)
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Close tears down the conn of HuloDaemon
func (term *HuloTerminal) Close() {
	cs := make(chan os.Signal, 1)
	signal.Notify(cs, os.Interrupt)
	<-cs
	term.Daemon.Close()
	os.Exit(0)
}
