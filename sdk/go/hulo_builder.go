package hulo

import (
	"errors"
	"flag"
	"fmt"
	"net"

	grpc "google.golang.org/grpc"
)

var port = flag.Int("p", 0, "")

func init() {
	flag.Parse()
	if *port == 0 {
		panic(errors.New("invalid port"))
	}
}

type Hulo struct {
	*HuloInterface
}

func NewHulo() *Hulo {
	return &Hulo{
		HuloInterface: NewHuloInterface(),
	}
}

func (hulo *Hulo) Completion(call HuloCompletion) {
	hulo.HuloInterface.completer = call
}

func (hulo *Hulo) Call(name string, call HuloCall) {
	hulo.HuloInterface.Register(name, call)
}

func (hulo *Hulo) Run() {
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	RegisterHuloInterfaceServer(gsrv, hulo.HuloInterface)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}

var HuloBuilder = NewHulo()
