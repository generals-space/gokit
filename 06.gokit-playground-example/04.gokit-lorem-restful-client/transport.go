package lorem_restful

import (
	"context"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	// ErrBadRouting url路径参数检测失败错误
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler 将endpoint类型的接口映射为常规http接口.
func MakeHTTPHandler(ctx context.Context, endpoint Endpoints) http.Handler {
	router := mux.NewRouter()
	routes := Initialize(router)

	loremHandler := httptransport.NewServer(
		endpoint.LoremEndpoint,
		DecodeLoremRequest,
		EncodeLoremResponse,
	)

	// mux提供的路由功能可以解析url中的路径参数, 且与标准库提供的handler兼容.
	routes.Lorem.Handler(loremHandler)

	return router
}
