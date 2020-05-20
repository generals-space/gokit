package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"gokit/pkg/model"
)

// User ...
type User struct {
	Name    string
	Title   string
	Company string
}

// UManagerServiceServer ...
type UManagerServiceServer struct{
	Users []*User
}

// GetUser ...
func (server *UManagerServiceServer) GetUser(ctx context.Context, req *model.GetUserRequest) (res *model.GetUserResponse, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			return &model.GetUserResponse{
				Name:    u.Name,
				Title:   u.Title,
				Company: u.Company,
			}, nil
		}
	}
	return nil, model.ErrUserNotFound
}

// SetTitle ...
func (server *UManagerServiceServer) SetTitle(ctx context.Context, req *model.SetTitleRequest) (res *model.Empty, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			u.Title = req.Title
			return &model.Empty{}, nil
		}
	}
	return &model.Empty{}, model.ErrUserNotFound
}

// Dispatch ...
func (server *UManagerServiceServer) Dispatch(ctx context.Context, req *model.DispatchRequest) (res *model.Empty, err error) {
	for _, u := range server.Users {
		if u.Name == req.Name {
			u.Company = req.Company
			return &model.Empty{}, nil
		}
	}
	return &model.Empty{}, model.ErrUserNotFound
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
	model.RegisterUserManagerServiceServer(rpcServer, uManagerServiceServer)
	reflection.Register(rpcServer)
	log.Println("server: 等待连接")
	if err := rpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
