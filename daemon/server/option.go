package server

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"go.uber.org/zap"
)

type DaemonOpt struct {
	Port   *int
	Logger *bool
}

func (opt DaemonOpt) Parse() {
	if len(os.Args) == 1 {
		opt.print()
		os.Exit(0)
	}
	flag.Parse()
	if *opt.Port == 0 {
		zap.S().Error(errors.New("invalid port"))
	}
}

func (DaemonOpt) print() {
	fmt.Println(`Usage:
  -p specify daemon port
  -l enable logger`)
}

var opt = DaemonOpt{
	Port:   flag.Int("p", 0, ""),
	Logger: flag.Bool("l", false, ""),
}
