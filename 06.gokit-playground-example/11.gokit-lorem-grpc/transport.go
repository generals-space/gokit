package lorem_grpc

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	"github.com/generals-space/gokit/06.gokit-playground-example/11.gokit-lorem-grpc/pb"
)

// LoremServiceServer 实现lorem.proto中的Lorem服务
type LoremServiceServer struct {
	lorem grpctransport.Handler
}

// Lorem 实现lorem.proto中声明的service接口
func (s *LoremServiceServer) Lorem(ctx context.Context, r *pb.LoremRequest) (*pb.LoremResponse, error) {
	_, resp, err := s.lorem.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.LoremResponse), nil
}

// NewGRPCServer ...
func NewGRPCServer(_ context.Context, endpoint Endpoints) pb.LoremServer {
	loremHandler := grpctransport.NewServer(
		endpoint.LoremEndpoint,
		DecodeGRPCLoremRequest,
		EncodeGRPCLoremResponse,
	)
	return &LoremServiceServer{
		lorem: loremHandler,
	}
}
