package main

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -L. -lhulo
#include "hulo.h"
*/
import "C"
import (
	context "context"
	"encoding/json"
	"fmt"
	"net"
	"unsafe"

	"github.com/ansurfen/cushion/cgo"
	"github.com/ansurfen/cushion/utils"
	grpc "google.golang.org/grpc"
)

type HuloInterface struct {
	UnimplementedHuloInterfaceServer
	dict map[string]func(string) string
}

func NewHuloInterface() *HuloInterface {
	return &HuloInterface{
		dict: make(map[string]func(string) string),
	}
}

func (hi *HuloInterface) Callback(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	if cb, ok := hi.dict[req.Func]; ok {
		res := cb(utils.JsonStr(utils.NewJsonObject(map[string]utils.JsonValue{
			"Func": utils.NewJsonString(req.Func),
			"Arg":  utils.NewJsonString(req.Arg),
		})))
		var ret CallResponse
		err := json.Unmarshal([]byte(res), &ret)
		return &ret, err
	}
	return &CallResponse{Buf: "unknown"}, nil
}

func (*HuloInterface) GetCandidate(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	return &CompletionResponse{}, nil
}

func (hi *HuloInterface) Register(name string, cb func(string) string) {
	hi.dict[name] = cb
}

type Hulo struct {
	hi *HuloInterface
}

//export newHulo
func newHulo() *C.Hulo {
	hulo := Hulo{
		hi: NewHuloInterface(),
	}
	ret := (*C.Hulo)(cgo.CMalloc(cgo.CSize_t(unsafe.Sizeof(C.Hulo{}))))
	ret.ptr = unsafe.Pointer(&hulo)
	return ret
}

//export registerCall
func registerCall(hulo *C.Hulo, name *C.char, cb C.Call) {
	h := cgo.CastPtr[Hulo](hulo.ptr)
	h.hi.dict[C.GoString(name)] = func(s string) string {
		str := C.huloCall(cb, C.CString(s))
		return C.GoString(str)
	}
}

//export run
func run(hulo *C.Hulo, port *C.char) {
	h := cgo.CastPtr[Hulo](hulo.ptr)
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", C.GoString(port)))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer()
	RegisterHuloInterfaceServer(gsrv, h.hi)
	if err := gsrv.Serve(listen); err != nil {
		panic(err)
	}
}

func main() {}
