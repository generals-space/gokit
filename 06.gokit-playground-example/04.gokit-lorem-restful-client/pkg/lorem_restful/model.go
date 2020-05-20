package lorem_restful

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// EncodeLoremRequest 将传入的LoremRequest对象转换成路径为/lorem/{type}/{min}/{max}的http.Request对象.
// 为了完成这个操作, 需要借助mux提供的路由工具.
// 注意这是一个闭包, 因为我们需要使用route参数构建url.
// 去主调函数看一下生成route的方法, 虽然可以在这个函数内部得到,
// 但是考虑到实际项目中路由接口众多, 为每个接口生成独立router会很烦琐.
func EncodeLoremRequest(route *mux.Route) httptransport.EncodeRequestFunc {
	return func(_ context.Context, r *http.Request, req interface{}) (err error) {
		loremRequest := req.(LoremRequest)

		r.URL, err = route.Host(r.URL.Host).URL(
			"type", loremRequest.RequestType,
			"min", fmt.Sprintf("%d", loremRequest.Min),
			"max", fmt.Sprintf("%d", loremRequest.Max),
		)
		if err != nil {
			return
		}

		methods, err := route.GetMethods()
		if err == nil {
			r.Method = methods[0]
		}
		return nil
	}
}

// DecodeLoremResponse 将服务端返回的json数据转换成LoremResponse对象
func DecodeLoremResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	loremResponse := &LoremResponse{}
	json.Unmarshal(body, loremResponse)
	return *loremResponse, nil
}

// DecodeLoremRequest 解析url路径参数为endpoint接口所需的LoremRequest对象
func DecodeLoremRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

// EncodeLoremResponse 将endpoint接口返回的响应转换成json数据
func EncodeLoremResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
