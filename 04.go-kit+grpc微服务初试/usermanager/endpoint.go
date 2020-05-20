package usermanager

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"gokit/common"
)

// Endpoint类型要求参数和返回值必须为interface类型, 而不能是具体类型.
// 所以makeXXXEndpoint只能返回如下的函数对象

func makeListEndpoint(srv *UserManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// req := request.(*common.Empty)
		users, err := srv.List()
		if err != nil {
			return nil, err
		}
		response := &common.UserList{}
		for _, user := range users {
			response.List = append(response.List, user)
		}
		return response, nil
	}
}

func makeAddUserEndpoint(srv *UserManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*common.UserList)
		return &common.Empty{}, srv.AddUser(req)
	}
}

func makeDispatchEndpoint(srv *UserManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*common.DispatchRequest)
		return &common.Empty{}, srv.Dispatch(req.Name, req.Company)
	}
}
