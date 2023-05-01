// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hulo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HuloInterfaceClient is the client API for HuloInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HuloInterfaceClient interface {
	Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error)
	Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type huloInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewHuloInterfaceClient(cc grpc.ClientConnInterface) HuloInterfaceClient {
	return &huloInterfaceClient{cc}
}

func (c *huloInterfaceClient) Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*CompletionResponse, error) {
	out := new(CompletionResponse)
	err := c.cc.Invoke(ctx, "/Hulo.HuloInterface/Completion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huloInterfaceClient) Call(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/Hulo.HuloInterface/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HuloInterfaceServer is the server API for HuloInterface service.
// All implementations must embed UnimplementedHuloInterfaceServer
// for forward compatibility
type HuloInterfaceServer interface {
	Completion(context.Context, *CompletionRequest) (*CompletionResponse, error)
	Call(context.Context, *CallRequest) (*CallResponse, error)
	mustEmbedUnimplementedHuloInterfaceServer()
}

// UnimplementedHuloInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedHuloInterfaceServer struct {
}

func (*UnimplementedHuloInterfaceServer) Completion(context.Context, *CompletionRequest) (*CompletionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Completion not implemented")
}
func (*UnimplementedHuloInterfaceServer) Call(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedHuloInterfaceServer) mustEmbedUnimplementedHuloInterfaceServer() {}

func RegisterHuloInterfaceServer(s *grpc.Server, srv HuloInterfaceServer) {
	s.RegisterService(&_HuloInterface_serviceDesc, srv)
}

func _HuloInterface_Completion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuloInterfaceServer).Completion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hulo.HuloInterface/Completion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuloInterfaceServer).Completion(ctx, req.(*CompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuloInterface_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuloInterfaceServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hulo.HuloInterface/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuloInterfaceServer).Call(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HuloInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Hulo.HuloInterface",
	HandlerType: (*HuloInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Completion",
			Handler:    _HuloInterface_Completion_Handler,
		},
		{
			MethodName: "Call",
			Handler:    _HuloInterface_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interface/hulo.proto",
}