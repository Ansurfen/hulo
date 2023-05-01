// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HuloDaemonClient is the client API for HuloDaemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HuloDaemonClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceResponse, error)
	CloseService(ctx context.Context, in *CloseServiceRequest, opts ...grpc.CallOption) (*CloseServiceResponse, error)
}

type huloDaemonClient struct {
	cc grpc.ClientConnInterface
}

func NewHuloDaemonClient(cc grpc.ClientConnInterface) HuloDaemonClient {
	return &huloDaemonClient{cc}
}

func (c *huloDaemonClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/proto.HuloDaemon/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huloDaemonClient) StartService(ctx context.Context, in *StartServiceRequest, opts ...grpc.CallOption) (*StartServiceResponse, error) {
	out := new(StartServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.HuloDaemon/StartService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *huloDaemonClient) CloseService(ctx context.Context, in *CloseServiceRequest, opts ...grpc.CallOption) (*CloseServiceResponse, error) {
	out := new(CloseServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.HuloDaemon/CloseService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HuloDaemonServer is the server API for HuloDaemon service.
// All implementations must embed UnimplementedHuloDaemonServer
// for forward compatibility
type HuloDaemonServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	StartService(context.Context, *StartServiceRequest) (*StartServiceResponse, error)
	CloseService(context.Context, *CloseServiceRequest) (*CloseServiceResponse, error)
	mustEmbedUnimplementedHuloDaemonServer()
}

// UnimplementedHuloDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedHuloDaemonServer struct {
}

func (*UnimplementedHuloDaemonServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedHuloDaemonServer) StartService(context.Context, *StartServiceRequest) (*StartServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartService not implemented")
}
func (*UnimplementedHuloDaemonServer) CloseService(context.Context, *CloseServiceRequest) (*CloseServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseService not implemented")
}
func (*UnimplementedHuloDaemonServer) mustEmbedUnimplementedHuloDaemonServer() {}

func RegisterHuloDaemonServer(s *grpc.Server, srv HuloDaemonServer) {
	s.RegisterService(&_HuloDaemon_serviceDesc, srv)
}

func _HuloDaemon_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuloDaemonServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HuloDaemon/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuloDaemonServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuloDaemon_StartService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuloDaemonServer).StartService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HuloDaemon/StartService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuloDaemonServer).StartService(ctx, req.(*StartServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HuloDaemon_CloseService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HuloDaemonServer).CloseService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.HuloDaemon/CloseService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HuloDaemonServer).CloseService(ctx, req.(*CloseServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HuloDaemon_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HuloDaemon",
	HandlerType: (*HuloDaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _HuloDaemon_Ping_Handler,
		},
		{
			MethodName: "StartService",
			Handler:    _HuloDaemon_StartService_Handler,
		},
		{
			MethodName: "CloseService",
			Handler:    _HuloDaemon_CloseService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "daemon/proto/daemon.proto",
}
