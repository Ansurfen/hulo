package terminal

import (
	"errors"
	"fmt"
	"strings"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type CommandArgv struct {
	cmd     string
	raw     string
	rawArgv string
	argv    []string
}

// Command is an interface to abstract any command to be parsed
type Command interface {
	// Exec to execute command
	Exec(*HuloTerminal, CommandArgv) (string, error)
}

type ExitCommand struct{}

func (ExitCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	return "", errors.New(Exit)
}

type PwdCommand struct{}

func (PwdCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	return term.IO.pw.Pwd(), nil
}

type CdCommand struct{}

func (CdCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	if len(arg.argv) > 0 {
		term.IO.pw.Walk(arg.argv[0])
	}
	return "", nil
}

type NoneCommand struct{}

func (NoneCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	return "", errors.New(NotFoundCmd)
}

type INTCommand struct{}

func (INTCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	return "", errors.New(ClearInterrupt)
}

type TestCommand struct{}

func (TestCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	return "", errors.New(SetInterrupt)
}

type EchoCommand struct{}

func (EchoCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	return "", nil
}

type LibCommand struct {
	fn func(term string)
}

func (c LibCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	c.fn(term.toString())
	return "", nil
}

type CallbackCommand struct {
	fn func(arg CommandArgv)
}

func (c CallbackCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	c.fn(arg)
	return "", nil
}

type LuaCommand struct {
	fn     *lua.LFunction
	loader string
}

func (c LuaCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	lvm := term.VM.Interp()
	lfunc := lvm.NewFunctionFromProto(c.fn.Proto)
	lvm.Push(lfunc)
	return "", term.VM.FastEvalFunc(lfunc,
		[]lua.LValue{luar.New(lvm, term), lua.LString(c.loader)})
}

type HuloCommand struct{}

func (HuloCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	switch a := ParseCommandArg(arg.rawArgv); a.cmd {
	case "load":
		return huloLoadCommand{}.Exec(term, a)
	case "install":
		return huloInstallCommand{}.Exec(term, a)
	case "lang":
		return huloLangCommand{}.Exec(term, a)
	case "daemon":
		return huloDaemonCommand{}.Exec(term, a)
	case "branch":
		return huloBranchCommand{}.Exec(term, a)
	default:
	}
	return "", nil
}

type huloBranchCommand struct{}

func (huloBranchCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	if len(arg.argv) == 0 {
		return "", nil
	}
	if _, ok := term.Completer.schemas[arg.argv[0]]; !ok {
		term.IO.WriteStderr(term.Patiner.errColor.Render("[ERROR] invalid schema"))
		return "", nil
	}
	term.Completer.SetSchema(arg.argv[0])
	cmd := term.Reg.Query(arg.argv[0])
	cmd.Exec(term, arg)
	return "", nil
}

type huloLoadCommand struct{}

func (huloLoadCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	if len(arg.argv) == 0 {
		return "", nil
	}
	term.Reg.LoadLoader(arg.argv[0])
	return "", nil
}

type huloInstallCommand struct{}

func (huloInstallCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	fmt.Println("install...")
	return "", nil
}

type huloLangCommand struct{}

func (huloLangCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	if len(arg.argv) == 0 {
		term.IO.WriteStderr(term.Patiner.errColor.Render("[ERROR] arg is null"))
		return "", nil
	}
	term.Env.Lang = arg.argv[0]
	return "", errors.New(Restart)
}

type huloDaemonCommand struct{}

func (huloDaemonCommand) Exec(term *HuloTerminal, arg CommandArgv) (string, error) {
	if len(arg.argv) == 0 {
		return "", nil
	}
	switch arg.argv[0] {
	case "ping":
		if err := term.Daemon.Ping(); err != nil {
			fmt.Println("fail to dial daemon")
		}
	}
	return "", nil
}

// ParseCommandArg parse command string to
func ParseCommandArg(cmd string) CommandArgv {
	argv := CommandArgv{
		raw: cmd,
	}
	res := strings.Split(strings.TrimSpace(cmd), " ")
	if len(res) > 0 {
		argv.cmd = res[0]
	}
	if len(res) > 1 {
		argv.argv = append(argv.argv, res[1:]...)
		argv.rawArgv = strings.Join(res[1:], " ")
	}
	return argv
}
