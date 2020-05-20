package main

import (
	"context"
	"fmt"
	"time"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"gokit/pb"
	"gokit/pkg/lorem_grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	loremEndpoint := grpctransport.NewClient(
		conn,
		"pb.Lorem", // 服务名称
		"Lorem",    // 方法名称
		lorem_grpc.EncodeGRPCLoremRequest,
		lorem_grpc.DecodeGRPCLoremResponse,
		pb.LoremResponse{},
	).Endpoint()
	endpoints := lorem_grpc.Endpoints{
		LoremEndpoint: loremEndpoint,
	}

	ctx := context.Background()
	msg, err := endpoints.Lorem(ctx, "Sentence", 5, 20)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(msg)
}

/*
// 常规的grpc客户端
func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	cli := pb.NewLoremClient(conn)
	fmt.Println("department manager service connected")
	ctx := context.Background()
	req := &pb.LoremRequest{
		RequestType: "Sentence",
		Min:         5,
		Max:         20,
	}
	resp, _ := cli.Lorem(ctx, req)
	fmt.Printf("%+v\n", resp)
}
*/
