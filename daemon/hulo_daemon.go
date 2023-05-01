package main

import (
	"hulo/daemon/server"

	"github.com/ansurfen/cushion/utils"
)

func init() {
	utils.InitLoggerWithDefault()
}

func main() {
	daemon := server.NewHuloDaemon()
	daemon.Run()
}
