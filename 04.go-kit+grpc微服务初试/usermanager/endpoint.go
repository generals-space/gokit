package usermanager

import (
	"context"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	"github.com/go-kit/kit/endpoint"
)

// Endpoint类型要求参数和返回值必须为interface类型, 而不能是具体类型.
// 所以makeXXXEndpoint只能返回如下的函数对象

func makeAddUserEndpoint(srv *UserManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(common.AddUserRequest)
		return &common.Empty{}, srv.AddUser(req.Name, req.Company)
	}
}

func makeGetUserEndpoint(srv *UserManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(common.GetUserRequest)
		user, err := srv.GetUser(req.Name)
		if err != nil {
			return nil, err
		}
		return &common.GetUserResponse{
			Name:    user.Name,
			Company: user.Company,
		}, nil
	}
}

func makeDispatchEndpoint(srv *UserManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(common.DispatchRequest)
		return &common.Empty{}, srv.Dispatch(req.Name, req.Company)
	}
}
