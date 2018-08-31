package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// UManagerServiceServer ...
type UManagerServiceServer struct{}

// GetUser ...
func (server *UManagerServiceServer) GetUser(ctx context.Context, req *GetUserRequest) (res *GetUserResponse, err error) {
	user, err := userManager.GetUser(req.Name)
	if err != nil {
		return
	}

	return &GetUserResponse{
		Name:    user.Name,
		Title:   user.Title,
		Company: user.Company,
	}, nil
}

// SetTitle ...
func (server *UManagerServiceServer) SetTitle(ctx context.Context, req *SetTitleRequest) (res *Empty, err error) {
	return &Empty{}, userManager.SetTitle(req.Name, req.Title)
}

// Dispatch ...
func (server *UManagerServiceServer) Dispatch(ctx context.Context, req *DispatchRequest) (res *Empty, err error) {
	return &Empty{}, userManager.Dispatch(req.Name, req.Company)
}

// NewServer ...
func NewServer() {
	log.Println("server: 启动监听")
	lis, err := net.Listen("tcp", ServerAddr)
	if err != nil {
		panic(err)
	}
	rpcServer := grpc.NewServer()
	log.Println("server: 注册服务")
	RegisterUserManagerServiceServer(rpcServer, &UManagerServiceServer{})
	reflection.Register(rpcServer)
	log.Println("server: 等待连接")
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
