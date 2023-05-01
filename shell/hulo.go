package shell

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

type HuloShell interface {
	StartProcess(path string, args string) error
	QueryProcessByName(name string) []Process
	StopProcess(pid int) error
}

func NewHuloShell() HuloShell {
	switch runtime.GOOS {
	case "windows":
		return PowerShell{}
	case "linux":
		return Bash{}
	case "darwin":
	}
	return nil
}

func Str2Float64(str string) float64 {
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		return f
	}
	return 0
}

func Str2Int(str string) int {
	if i, err := strconv.ParseInt(str, 10, 32); err == nil {
		return int(i)
	}
	return 0
}

type Process interface {
	Pid() int
}

func ELF(elf string) string {
	switch runtime.GOOS {
	case "windows":
		if !strings.HasSuffix(elf, ".exe") {
			return fmt.Sprintf("%s.exe", strings.TrimSpace(elf))
		}
	case "linux":
	case "darwin":
	default:
	}
	return elf
}
