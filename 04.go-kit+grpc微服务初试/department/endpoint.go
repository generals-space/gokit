package department

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)

func makeListEndpoint(srv *DepartmentManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		departments, err := srv.List()
		if err != nil {
			return nil, err
		}
		response := &common.DepartmentList{}
		for _, department := range departments {
			response.List = append(response.List, department)
		}
		return response, nil
	}
}

func makeCreateEndpoint(srv *DepartmentManager) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*common.Department)
		return &common.Empty{}, srv.Create(req)
	}
}
