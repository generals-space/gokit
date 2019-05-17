package lorem_consul

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// LoremRequest ...
type LoremRequest struct {
	RequestType string `json:"requestType"`
	Min         int    `json:"min"`
	Max         int    `json:"max"`
}

// LoremResponse ...
type LoremResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

// HealthRequest ...
type HealthRequest struct {
}

// HealthResponse ...
type HealthResponse struct {
	Status bool `json:"status"`
}

// Endpoints endpoint集合, 类似于常规http服务同一url前缀下提供的不同的子接口, 挂载路由时可以统一挂载.
type Endpoints struct {
	LoremEndpoint  endpoint.Endpoint
	HealthEndpoint endpoint.Endpoint
}

// MakeLoremEndpoint 将Lorem业务逻辑包装成一个endpoint.
func MakeLoremEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		min, max := int(req.Min), int(req.Max)
		txt, err := svc.Lorem(req.RequestType, min, max)

		if err != nil {
			return nil, err
		}

		return LoremResponse{Message: txt}, nil
	}

}

// MakeHealthEndpoint 将HealthCheck业务逻辑包装成endpoint
// 其实如果是直接返回true的话, 可以在这一步返回.
// 但是实际上健康检查服务可能需要检测数据库连接, 依赖的子服务状态等以确认服务是否真的可用.
func MakeHealthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		status := svc.HealthCheck()
		return HealthResponse{Status: status}, nil
	}
}
