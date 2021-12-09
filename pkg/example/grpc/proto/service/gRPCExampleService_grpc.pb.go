// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	message "pkg/example/grpc/proto/message"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GRPCExampleServiceClient is the client API for GRPCExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCExampleServiceClient interface {
	SayHello(ctx context.Context, in *message.HelloReq, opts ...grpc.CallOption) (*message.HelloReply, error)
}

type gRPCExampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCExampleServiceClient(cc grpc.ClientConnInterface) GRPCExampleServiceClient {
	return &gRPCExampleServiceClient{cc}
}

func (c *gRPCExampleServiceClient) SayHello(ctx context.Context, in *message.HelloReq, opts ...grpc.CallOption) (*message.HelloReply, error) {
	out := new(message.HelloReply)
	err := c.cc.Invoke(ctx, "/service.gRPCExampleService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCExampleServiceServer is the server API for GRPCExampleService service.
// All implementations must embed UnimplementedGRPCExampleServiceServer
// for forward compatibility
type GRPCExampleServiceServer interface {
	SayHello(context.Context, *message.HelloReq) (*message.HelloReply, error)
	mustEmbedUnimplementedGRPCExampleServiceServer()
}

// UnimplementedGRPCExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCExampleServiceServer struct {
}

func (UnimplementedGRPCExampleServiceServer) SayHello(context.Context, *message.HelloReq) (*message.HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGRPCExampleServiceServer) mustEmbedUnimplementedGRPCExampleServiceServer() {}

// UnsafeGRPCExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCExampleServiceServer will
// result in compilation errors.
type UnsafeGRPCExampleServiceServer interface {
	mustEmbedUnimplementedGRPCExampleServiceServer()
}

func RegisterGRPCExampleServiceServer(s grpc.ServiceRegistrar, srv GRPCExampleServiceServer) {
	s.RegisterService(&GRPCExampleService_ServiceDesc, srv)
}

func _GRPCExampleService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCExampleServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.gRPCExampleService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCExampleServiceServer).SayHello(ctx, req.(*message.HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCExampleService_ServiceDesc is the grpc.ServiceDesc for GRPCExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.gRPCExampleService",
	HandlerType: (*GRPCExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GRPCExampleService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gRPCExampleService.proto",
}