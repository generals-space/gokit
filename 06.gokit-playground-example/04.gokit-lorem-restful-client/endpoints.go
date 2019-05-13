package lorem_restful

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// LoremRequest ...
type LoremRequest struct {
	RequestType string
	Min         int
	Max         int
}

// LoremResponse ...
type LoremResponse struct {
	Message string `json:"message"`
	Err     error  `json:"err,omitempty"`
}

// Endpoints endpoint集合, 类似于常规http服务同一url前缀下提供的不同的子接口, 挂载路由时可以统一挂载.
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// MakeLoremEndpoint 将Lorem业务逻辑包装成一个endpoint.
// 虽然业务逻辑提供了3个方法, 但只有一个endpoint接口.
// 其实就算是常规http服务, 也只会提供一个接口`/lorem/{type}/{min}/{max}`, 只是路径参数而已.
// 可能是为了业务服务为了能更灵活, 所以对type的判断写在了这里而不是业务逻辑中.
// 注意: endpoint的规定返回的数据为一个interface{}, 和一个error.
// 当error为nil, 即业务逻辑正常时会调用encodeResponse方法
// 当error不为nil的时候, 默认会不经过encodeResponse直接返回error信息,
// 如果在transport中定义了ServerErrorEncoder方法则会调用其中指定的方法.
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

// Lorem ...
func (e Endpoints) Lorem(ctx context.Context, requestType string, min, max int) (string, error) {
	req := LoremRequest{
		RequestType: requestType,
		Min:         min,
		Max:         max,
	}
	resp, err := e.LoremEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	loremResp := resp.(LoremResponse)
	if loremResp.Err != nil {
		return "", errors.New(loremResp.Err.Error())
	}
	return loremResp.Message, nil
}
