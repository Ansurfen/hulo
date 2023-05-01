package terminal

import (
	"os"
	"path"
	"strings"
)

type PathWalk struct {
	Global bool
	Root   string
}

func NewPathWalk(g bool, spec ...string) *PathWalk {
	pw := &PathWalk{
		Global: g,
	}
	if !g {
		if root, err := os.Getwd(); err != nil {
			panic(err)
		} else {
			pw.Root = root
		}
	}
	if len(spec) > 0 {
		if os.Chdir(spec[0]) != nil {
			panic("")
		}
		pw.Root = spec[0]
	}
	return pw
}

func (pw *PathWalk) Walk(dir string) error {
	if !pw.Global {
		if !strings.HasPrefix(pw.Root, path.Join(pw.Pwd(), dir)) {
			return nil
		}
	}
	return os.Chdir(dir)
}

func (pw *PathWalk) Pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	if !pw.Global && pw.Root == pwd {
		return "/"
	}
	return pwd
}
