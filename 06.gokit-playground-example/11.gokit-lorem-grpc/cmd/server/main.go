package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"gokit/pb"
	"gokit/pkg/lorem_grpc"
)

func main() {
	// 创建业务服务
	var svc lorem_grpc.Service
	svc = lorem_grpc.LoremService{}
	endpoints := lorem_grpc.Endpoints{
		LoremEndpoint: lorem_grpc.MakeLoremEndpoint(svc),
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	handler := lorem_grpc.NewGRPCServer(ctx, endpoints)
	gRPCServer := grpc.NewServer()
	pb.RegisterLoremServer(gRPCServer, handler)

	fmt.Println("Starting server")
	fmt.Println(gRPCServer.Serve(listener))
}
