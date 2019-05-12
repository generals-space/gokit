package lorem_restful

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	// ErrBadRouting url路径参数检测失败错误
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// 解析url路径参数为endpoint接口所需的LoremRequest对象
func decodeLoremRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrBadRouting
	}

	vmin, ok := vars["min"]
	if !ok {
		return nil, ErrBadRouting
	}

	vmax, ok := vars["max"]
	if !ok {
		return nil, ErrBadRouting
	}

	min, _ := strconv.Atoi(vmin)
	max, _ := strconv.Atoi(vmax)
	return LoremRequest{
		RequestType: requestType,
		Min:         min,
		Max:         max,
	}, nil
}

// 将endpoint接口返回的响应转换成json数据
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// MakeHTTPHandler 将endpoint类型的接口映射为常规http接口.
func MakeHTTPHandler(ctx context.Context, endpoint Endpoints) http.Handler {
	handler := httptransport.NewServer(
		endpoint.LoremEndpoint,
		decodeLoremRequest,
		encodeResponse,
	)

	// mux提供的路由功能可以解析url中的路径参数, 且与标准库提供的handler兼容.
	r := mux.NewRouter()
	r.Methods("POST").Path("/lorem/{type}/{min}/{max}").Handler(handler)

	return r
}
