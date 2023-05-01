package hulo

import context "context"

var _ HuloInterfaceServer = &HuloInterface{}

type (
	HuloCall       func(req *CallRequest) (*CallResponse, error)
	HuloCompletion func(req *CompletionRequest) (*CompletionResponse, error)
)

type HuloInterface struct {
	UnimplementedHuloInterfaceServer
	dict      map[string]HuloCall
	completer HuloCompletion
}

func NewHuloInterface() *HuloInterface {
	return &HuloInterface{
		dict: make(map[string]HuloCall),
		completer: func(req *CompletionRequest) (*CompletionResponse, error) {
			return &CompletionResponse{}, nil
		},
	}
}

func (hi *HuloInterface) Call(ctx context.Context, req *CallRequest) (*CallResponse, error) {
	if call, ok := hi.dict[req.Func]; ok {
		return call(req)
	}
	return &CallResponse{Buf: "unknown"}, nil
}

func (hi *HuloInterface) Completion(ctx context.Context, req *CompletionRequest) (*CompletionResponse, error) {
	return hi.completer(req)
}

func (hi *HuloInterface) Register(name string, call HuloCall) {
	hi.dict[name] = call
}

func (hi *HuloInterface) UnRegister(name string) {
	delete(hi.dict, name)
}
