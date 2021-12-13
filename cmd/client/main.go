package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"pkg/example/grpc/proto/message"
	"pkg/example/grpc/proto/service"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:8848", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("connect faild: %v", err)
	}
	defer conn.Close()

	c := service.NewGRPCExampleServiceClient(conn)
	// 调用SayHello
	r, err := c.SayHello(context.Background(), &message.HelloReq{Context: "tim"})
	if err != nil {
		fmt.Printf("sayHello failed: %v", err)
	}
	if r != nil {
		fmt.Printf("SayHello: %s \n", r.Callback)
	}
}
