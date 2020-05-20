package department

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	transport_grpc "github.com/go-kit/kit/transport/grpc"

	"gokit/common"
)

// DManagerServiceServer ...
type DManagerServiceServer struct {
	CreateHandler transport_grpc.Handler
	ListHandler   transport_grpc.Handler
	PersonnelChangeHandler transport_grpc.Handler
}

// List ...
func (server *DManagerServiceServer) List(ctx context.Context, req *common.Empty) (res *common.DepartmentList, err error) {
	_, resp, err := server.ListHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*common.DepartmentList), nil
}

// Create ...
func (server *DManagerServiceServer) Create(ctx context.Context, req *common.Department) (res *common.Empty, err error) {
	_, resp, err := server.CreateHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*common.Empty), nil
}

// PersonnelChange ...
func (server *DManagerServiceServer) PersonnelChange(ctx context.Context, req *common.PersonnelChangeRequest)(res *common.Empty, err error){
	_, resp, err := server.PersonnelChangeHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*common.Empty), nil
}

func decodeGrpcRequest(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

func encodeGrpcResponse(_ context.Context, req interface{}) (interface{}, error) {
	return req, nil
}

// NewGrpcServer ...
func NewGrpcServer(srv *DepartmentManager) *DManagerServiceServer {
	listHandler := transport_grpc.NewServer(
		makeListEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	createHandler := transport_grpc.NewServer(
		makeCreateEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	personnelChangeHandler := transport_grpc.NewServer(
		makePersonnelChangeEndpoint(srv),
		decodeGrpcRequest,
		encodeGrpcResponse,
	)
	return &DManagerServiceServer{
		CreateHandler: createHandler,
		ListHandler:   listHandler,
		PersonnelChangeHandler: personnelChangeHandler,
	}
}

// StartGrpcTransport ...
func StartGrpcTransport(srv *DepartmentManager) {
	log.Println("starting department manager grpc transport...")

	lis, _ := net.Listen("tcp", common.DepartmentGrpcTransportAddr)
	gprcServer := grpc.NewServer()
	dManagerServiceServer := NewGrpcServer(srv)
	common.RegisterDepartmentManagerServiceServer(gprcServer, dManagerServiceServer)
	gprcServer.Serve(lis)
}
