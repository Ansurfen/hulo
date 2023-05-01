package terminal

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/ansurfen/cushion/cgo"

	"github.com/ansurfen/cushion/utils"
)

// HuloRegistry manage commands and loaders
type HuloRegistry struct {
	cmds    map[string]Command
	loaders map[string]*Loader
	term    *HuloTerminal
}

func NewHuloRegistry() *HuloRegistry {
	reg := &HuloRegistry{
		cmds:    make(map[string]Command),
		loaders: make(map[string]*Loader),
	}
	reg.cmds = map[string]Command{
		"cd":   CdCommand{},
		"pwd":  PwdCommand{},
		"int":  INTCommand{},
		"open": TestCommand{},
		"echo": EchoCommand{},
		"hulo": HuloCommand{},
	}
	loaders, err := os.ReadDir(LoaderPath)
	if err != nil {
		panic(err)
	}
	for _, loader := range loaders {
		if !loader.IsDir() {
			reg.loaders[utils.Filename(loader.Name())] = &Loader{
				name:  loader.Name(),
				state: LoaderSuspend,
			}
		}
	}
	return reg
}

// Query returns specify command to be registered in registry
// and return NoneCommand when command isn't exist
func (reg *HuloRegistry) Query(cmd string) Command {
	if c, ok := reg.cmds[cmd]; ok {
		return c
	}
	return NoneCommand{}
}

// RegisterCmd to add command to registry
func (reg *HuloRegistry) RegisterCmd(name string, cmd Command) {
	if _, ok := reg.cmds[name]; ok {
		reg.term.IO.WriteStderr("\n" + reg.term.Patiner.errColor.Render(fmt.Sprintf("[ERROR] %s plugin exist already", name)))
	} else {
		reg.cmds[name] = cmd
	}
}

// UnregisterCmd to remove command from registry
func (reg *HuloRegistry) UnregisterCmd(name string) {
	delete(reg.cmds, name)
}

// LoadLoader to load specify loader
func (reg *HuloRegistry) LoadLoader(name string) error {
	if loader, ok := reg.loaders[name]; ok {
		if loader.state == LoaderStop {
			return errors.New(FailLoadLoader)
		}
		err := reg.term.VM.EvalFile(path.Join(LoaderPath, name+".lua"))
		if err != nil {
			loader.state = LoaderStop
			return errors.New(FailLoadLoader)
		}
		loader.state = LoaderReady
	}
	return nil
}

// LoadLib to load specify dynamic library (so, dylib, dll)
func (reg *HuloRegistry) LoadLib(name string) {
	p, _ := utils.NewPlugin(path.Join(LibPath, name))
	reg.RegisterCmd(name, LibCommand{
		fn: func(term string) {
			grep, _ := p.Func("Exec")
			grep.Call(cgo.CastVoidPtr(cgo.CStr(term)))
		},
	})
}

// LoadComplete to load static completion rules
func (reg *HuloRegistry) LoadComplete(name string) Exception {
	plugins, err := os.ReadDir(path.Join(CompletePath, name))
	if err != nil {
		return FailLoadComplete
	}
	for _, plugin := range plugins {
		if !plugin.IsDir() && plugin.Name() == reg.term.Env.Lang+".lua" {
			reg.term.VM.EvalFile(path.Join(CompletePath, name, plugin.Name()))
		}
	}
	return NoException
}

// SetLoaderRunning to set the state of loader to LoaderRunning
func (reg *HuloRegistry) SetLoaderRunning(name string) error {
	return reg.setLoaderState(name, RUNNING)
}

// SetLoaderReady to set the state of loader to LoaderReady
func (reg *HuloRegistry) SetLoaderReady(name string) error {
	return reg.setLoaderState(name, READY)
}

// SetLoaderStop to set the state of loader to LoaderStop
func (reg *HuloRegistry) SetLoaderStop(name string) error {
	return reg.setLoaderState(name, STOP)
}

// setLoaderState could set the any state of loader
func (reg *HuloRegistry) setLoaderState(name string, state int) error {
	if loader, ok := reg.loaders[name]; ok {
		loader.state = state
		return nil
	}
	return errors.New("invalid loader")
}

// LoaderState returns state of specify loader
func (reg *HuloRegistry) LoaderState(name string) int {
	if loader, ok := reg.loaders[name]; ok {
		return loader.state
	}
	return LoaderUndefined
}

const (
	LoaderSuspend = iota
	LoaderReady
	LoaderRunning
	LoaderStop
	LoaderUndefined
)

type Loader struct {
	name  string
	state int
}
