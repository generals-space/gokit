package lorem_restful

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

var (
	ErrRequestTypeNotFound = errors.New("Request type only valid for word, sentence and paragraph")
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

		var (
			txt      string
			min, max int
		)

		min, max = req.Min, req.Max

		if strings.EqualFold(req.RequestType, "Word") {
			txt = svc.Word(min, max)
		} else if strings.EqualFold(req.RequestType, "Sentence") {
			txt = svc.Sentence(min, max)
		} else if strings.EqualFold(req.RequestType, "Paragraph") {
			txt = svc.Paragraph(min, max)
		} else {
			return nil, ErrRequestTypeNotFound
		}
		// return LoremResponse{Message: txt}, nil
		return LoremResponse{Message: txt}, errors.New("test error")
	}
}
