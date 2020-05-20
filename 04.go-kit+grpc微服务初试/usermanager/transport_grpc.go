package usermanager

import (
	"context"
	"log"
	"net"

	transport_grpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"gokit/common"
)

// UManagerServiceServer ...
type UManagerServiceServer struct {
	ListHandler  transport_grpc.Handler
	AddUserHandler  transport_grpc.Handler
	DispatchHandler transport_grpc.Handler
}

// List ...
func (server *UManagerServiceServer) List(ctx context.Context, req *common.Empty) (res *common.UserList, err error) {
	_, rsp, err := server.ListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*common.UserList), nil
}

// AddUser ...
func (server *UManagerServiceServer) AddUser(ctx context.Context, req *common.UserList) (res *common.Empty, err error) {
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
	log.Println("starting user manager grpc transport...")

	ListHandler := transport_grpc.NewServer(
		makeListEndpoint(service),
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
		ListHandler:  ListHandler,
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
