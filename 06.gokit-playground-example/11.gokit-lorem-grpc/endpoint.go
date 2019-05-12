package lorem_grpc

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

// LoremRequest ...
type LoremRequest struct {
	RequestType string
	Min         int32
	Max         int32
}

// LoremResponse ...
type LoremResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}

// Endpoints ...
// 在restful示例中Endpoint是同一url前缀的路由集合, 但是url不同子路径是挂载到不同endpoint成员的.
// 在grpc示例中同样是一组接口的集合, 但是对外提供服务暴露的是Endpoints本身.
// 客户端想要调用不同的endpoint成员, 就要调用Endpoints结构中的不同成员方法.
type Endpoints struct {
	LoremEndpoint endpoint.Endpoint
}

// MakeLoremEndpoint 创建LoremEndpoint端点
func MakeLoremEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoremRequest)

		min, max := int(req.Min), int(req.Max)
		txt, err := svc.Lorem(ctx, req.RequestType, min, max)

		if err != nil {
			return nil, err
		}

		return LoremResponse{Message: txt}, nil
	}

}

// Lorem 将Endpoints中的成员端点包装成方法, 这样grpc客户端能直接调用而不是通过endpoint成员对象调用
func (e Endpoints) Lorem(ctx context.Context, requestType string, min, max int) (string, error) {
	req := LoremRequest{
		RequestType: requestType,
		Min:         int32(min),
		Max:         int32(max),
	}
	resp, err := e.LoremEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	loremResp := resp.(LoremResponse)
	if loremResp.Err != "" {
		return "", errors.New(loremResp.Err)
	}
	return loremResp.Message, nil
}
