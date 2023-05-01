package terminal

import (
	"fmt"
	"hulo/shell"
	"path"
)

const (
	HSMountPlugin = iota
	HSLoadTheme
	HSInitDaemon
)

// HuloScheduler, is used to lazy hot reload and initialize
var HuloScheduler map[uint8]func(*HuloTerminal) error

func hsMountPlugin(term *HuloTerminal) error {
	for _, loader := range term.Env.Boot {
		if err := term.Reg.LoadLoader(loader); err != nil {
			term.IO.WriteStdout(term.Logger.Error("fail to boot loader"))
		}
	}
	term.VM.EvalFile("sdk.lua")
	return nil
}

func hsLoadTheme(term *HuloTerminal) error {
	term.Patiner.LoadTheme(term.Env.Theme)
	return nil
}

func hsInitDaemon(term *HuloTerminal) error {
	sh := shell.NewHuloShell()
	sh.StartProcess(path.Join(BinPath, shell.ELF("hulo_daemon")), fmt.Sprintf("-p %d", term.Env.Daemon))
	term.Daemon = NewDaemon(term.Env.Daemon)
	if err := term.Daemon.Ping(); err != nil {
		term.IO.WriteStderr(term.Logger.Error(FailDailDaemon))
	}
	term.Daemon.term = term
	return nil
}
