package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"pkg/example/grpc/proto/message"
	"pkg/example/grpc/proto/service"
)

type ExampleGrpcImpl struct {
	service.UnimplementedGRPCExampleServiceServer
}

func (e *ExampleGrpcImpl) SayHello(ctx context.Context, req *message.HelloReq) (*message.HelloReply, error) {
	fmt.Printf("rpc is done:%v", req.Context)
	result := &message.HelloReply{Callback: "is ok"}
	return result, nil
}

func main() {
	// 监听本地的8848端口
	lis, err := net.Listen("tcp", "localhost:8848")
	if err != nil {
		fmt.Printf("listen failed: %v", err)
		return
	}
	s := grpc.NewServer()                                           // 创建gRPC服务器
	service.RegisterGRPCExampleServiceServer(s, &ExampleGrpcImpl{}) // 在gRPC服务端注册服务

	reflection.Register(s) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
	fmt.Printf("listen success: %v", "grpc")
}
