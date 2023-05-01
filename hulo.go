package main

import (
	"hulo/terminal"
)

func main() {
	hulo := terminal.NewHulo()
	go hulo.Close()
ljmp:
	hulo.Initialize()
	for {
		// checkpoint: handle event before read
		hulo.EventBus.BeforeRead()
		// accept input from IO
		in := hulo.Read()
		// checkpoint: handle event after read
		hulo.EventBus.AfterRead()
		argv := terminal.ParseCommandArg(in)
		// checkpoint:
		hulo.EventBus.BeforeExec()
		sig := hulo.Exec(argv)
		// checkpoint:
		hulo.EventBus.AfterExec()
		switch sig.Op {
		// hot reload from head of code, radically
		case terminal.LJMP:
			goto ljmp
		// make use of eventbus to simulate hot reload
		case terminal.SJMP:
			// it'll be hot reload in BeforeRead
		case terminal.CLI:
			// clear interpreter to close IO
		case terminal.STI:
			// set interpreter to open IO
		case terminal.NOP: // do nothing
			fallthrough
		default:
		}
	}
}
