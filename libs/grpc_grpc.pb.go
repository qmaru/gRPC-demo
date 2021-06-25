// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// HelloGRPCClient is the client API for HelloGRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloGRPCClient interface {
	// Checking Server
	ServerCheck(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type helloGRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloGRPCClient(cc grpc.ClientConnInterface) HelloGRPCClient {
	return &helloGRPCClient{cc}
}

func (c *helloGRPCClient) ServerCheck(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/pb.HelloGRPC/ServerCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloGRPCServer is the server API for HelloGRPC service.
// All implementations must embed UnimplementedHelloGRPCServer
// for forward compatibility
type HelloGRPCServer interface {
	// Checking Server
	ServerCheck(context.Context, *Ping) (*Pong, error)
	mustEmbedUnimplementedHelloGRPCServer()
}

// UnimplementedHelloGRPCServer must be embedded to have forward compatible implementations.
type UnimplementedHelloGRPCServer struct {
}

func (UnimplementedHelloGRPCServer) ServerCheck(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerCheck not implemented")
}
func (UnimplementedHelloGRPCServer) mustEmbedUnimplementedHelloGRPCServer() {}

// UnsafeHelloGRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloGRPCServer will
// result in compilation errors.
type UnsafeHelloGRPCServer interface {
	mustEmbedUnimplementedHelloGRPCServer()
}

func RegisterHelloGRPCServer(s grpc.ServiceRegistrar, srv HelloGRPCServer) {
	s.RegisterService(&HelloGRPC_ServiceDesc, srv)
}

func _HelloGRPC_ServerCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloGRPCServer).ServerCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HelloGRPC/ServerCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloGRPCServer).ServerCheck(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

// HelloGRPC_ServiceDesc is the grpc.ServiceDesc for HelloGRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloGRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HelloGRPC",
	HandlerType: (*HelloGRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerCheck",
			Handler:    _HelloGRPC_ServerCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}