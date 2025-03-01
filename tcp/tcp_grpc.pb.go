// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: tcp/tcp.proto

package tcp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TcpMessagingClient is the client API for TcpMessaging service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TcpMessagingClient interface {
	SendMessage(ctx context.Context, in *Tcp, opts ...grpc.CallOption) (*Tcp, error)
}

type tcpMessagingClient struct {
	cc grpc.ClientConnInterface
}

func NewTcpMessagingClient(cc grpc.ClientConnInterface) TcpMessagingClient {
	return &tcpMessagingClient{cc}
}

func (c *tcpMessagingClient) SendMessage(ctx context.Context, in *Tcp, opts ...grpc.CallOption) (*Tcp, error) {
	out := new(Tcp)
	err := c.cc.Invoke(ctx, "/tcp.tcpMessaging/sendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TcpMessagingServer is the server API for TcpMessaging service.
// All implementations must embed UnimplementedTcpMessagingServer
// for forward compatibility
type TcpMessagingServer interface {
	SendMessage(context.Context, *Tcp) (*Tcp, error)
	mustEmbedUnimplementedTcpMessagingServer()
}

// UnimplementedTcpMessagingServer must be embedded to have forward compatible implementations.
type UnimplementedTcpMessagingServer struct {
}

func (UnimplementedTcpMessagingServer) SendMessage(context.Context, *Tcp) (*Tcp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedTcpMessagingServer) mustEmbedUnimplementedTcpMessagingServer() {}

// UnsafeTcpMessagingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TcpMessagingServer will
// result in compilation errors.
type UnsafeTcpMessagingServer interface {
	mustEmbedUnimplementedTcpMessagingServer()
}

func RegisterTcpMessagingServer(s grpc.ServiceRegistrar, srv TcpMessagingServer) {
	s.RegisterService(&TcpMessaging_ServiceDesc, srv)
}

func _TcpMessaging_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tcp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TcpMessagingServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tcp.tcpMessaging/sendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TcpMessagingServer).SendMessage(ctx, req.(*Tcp))
	}
	return interceptor(ctx, in, info, handler)
}

// TcpMessaging_ServiceDesc is the grpc.ServiceDesc for TcpMessaging service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TcpMessaging_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tcp.tcpMessaging",
	HandlerType: (*TcpMessagingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sendMessage",
			Handler:    _TcpMessaging_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tcp/tcp.proto",
}
