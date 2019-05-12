package lorem_grpc

import (
	"context"

	"github.com/generals-space/gokit/06.gokit-playground-example/11.gokit-lorem-grpc/pb"
)

// EncodeGRPCLoremRequest ...
func EncodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(LoremRequest)
	return &pb.LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// DecodeGRPCLoremRequest ...
func DecodeGRPCLoremRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.LoremRequest)
	return LoremRequest{
		RequestType: req.RequestType,
		Max:         req.Max,
		Min:         req.Min,
	}, nil
}

// EncodeGRPCLoremResponse ...
func EncodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(LoremResponse)
	return &pb.LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}

// DecodeGRPCLoremResponse ...
func DecodeGRPCLoremResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.LoremResponse)
	return LoremResponse{
		Message: resp.Message,
		Err:     resp.Err,
	}, nil
}
