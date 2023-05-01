package hulo

import "testing"

func TestHulo(t *testing.T) {
	HuloBuilder.Completion(func(req *CompletionRequest) (*CompletionResponse, error) {
		return &CompletionResponse{}, nil
	})
	HuloBuilder.Call("SayHello", func(req *CallRequest) (*CallResponse, error) {
		return &CallResponse{Buf: "Hello, I'm golang"}, nil
	})
	HuloBuilder.Run()
}
