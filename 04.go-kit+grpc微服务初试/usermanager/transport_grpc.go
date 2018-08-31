package usermanager

import (
	"context"
	"net"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// UManagerServiceServer ...
type UManagerServiceServer struct {
	GetUserHandler  transport_grpc.Handler
	AddUserHandler  transport_grpc.Handler
	DispatchHandler transport_grpc.Handler
}

// GetUser ...
func (server *UManagerServiceServer) GetUser(ctx context.Context, req *common.GetUserRequest) (res *common.GetUserResponse, err error) {
	_, rsp, err := server.GetUserHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*common.GetUserResponse), nil
}

// AddUser ...
func (server *UManagerServiceServer) AddUser(ctx context.Context, req *common.AddUserRequest) (res *common.Empty, err error) {
	_, rsp, err := server.AddUserHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*common.Empty), nil
}

// Dispatch ...
func (server *UManagerServiceServer) Dispatch(ctx context.Context, req *common.DispatchRequest) (res *common.Empty, err error) {
	_, rsp, err := server.DispatchHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*common.Empty), nil
}

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// NewGrpcServer ...
func NewGrpcServer(service *UserManager) *UManagerServiceServer {
	GetUserHandler := transport_grpc.NewServer(
		makeGetUserEndpoint(service),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	AddUserHandler := transport_grpc.NewServer(
		makeAddUserEndpoint(service),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	DispatchHandler := transport_grpc.NewServer(
		makeDispatchEndpoint(service),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)

	return &UManagerServiceServer{
		GetUserHandler:  GetUserHandler,
		AddUserHandler:  AddUserHandler,
		DispatchHandler: DispatchHandler,
	}
}

// StartGrpcTransport 启动http transport
func StartGrpcTransport(srv *UserManager) {

	lis, _ := net.Listen("tcp", common.UserManagerGrpcTransportAddr)
	gprcServer := grpc.NewServer()
	uManagerServiceServer := NewGrpcServer(srv)
	common.RegisterUserManagerServiceServer(gprcServer, uManagerServiceServer)
	gprcServer.Serve(lis)
}
