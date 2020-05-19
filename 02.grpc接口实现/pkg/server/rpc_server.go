package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gokit/pkg/model"
)

// UManagerServiceServer ...
type UManagerServiceServer struct{}

// GetUser ...
func (server *UManagerServiceServer) GetUser(ctx context.Context, req *model.GetUserRequest) (res *model.GetUserResponse, err error) {
	user, err := model.UM.GetUser(req.Name)
	if err != nil {
		return
	}

	return &model.GetUserResponse{
		Name:    user.Name,
		Title:   user.Title,
		Company: user.Company,
	}, nil
}

// SetTitle ...
func (server *UManagerServiceServer) SetTitle(ctx context.Context, req *model.SetTitleRequest) (res *model.Empty, err error) {
	return &model.Empty{}, model.UM.SetTitle(req.Name, req.Title)
}

// Dispatch ...
func (server *UManagerServiceServer) Dispatch(ctx context.Context, req *model.DispatchRequest) (res *model.Empty, err error) {
	return &model.Empty{}, model.UM.Dispatch(req.Name, req.Company)
}

// NewServer ...
func NewServer() {
	log.Println("server: 启动监听")
	lis, err := net.Listen("tcp", model.ServerAddr)
	if err != nil {
		panic(err)
	}
	rpcServer := grpc.NewServer()
	log.Println("server: 注册服务")
	model.RegisterUserManagerServiceServer(rpcServer, &UManagerServiceServer{})
	reflection.Register(rpcServer)
	log.Println("server: 等待连接")
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
