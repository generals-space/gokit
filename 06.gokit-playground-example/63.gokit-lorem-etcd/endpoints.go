package lorem_etcd

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

// Endpoints endpoint集合, 类似于常规http服务同一url前缀下提供的不同的子接口, 挂载路由时可以统一挂载.
type Endpoints struct {
	LoremEndpoint  endpoint.Endpoint
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
