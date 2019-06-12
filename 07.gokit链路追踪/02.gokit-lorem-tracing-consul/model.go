package lorem_tracing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DecodeLoremRequest ...
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

	request := LoremRequest{
		RequestType: requestType,
		Min:         min,
		Max:         max,
	}
	return request, nil
}

// DecodeHealthRequest ...
func DecodeHealthRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return HealthRequest{}, nil
}

// EncodeResponse 这是一个通用方法, 将对象转换成json字符串就可以了, 不用在乎对象类型.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	fmt.Println("iiiiiiiiiiiiiiiii")
	fmt.Printf("%+v\n", response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// EncodeLoremRequest ...
func EncodeLoremRequest(_ context.Context, req *http.Request, request interface{}) error {
	lr := request.(LoremRequest)
	p := "/" + lr.RequestType + "/" + strconv.Itoa(lr.Min) + "/" + strconv.Itoa(lr.Max)
	req.URL.Path += p
	return nil
}

// DecodeLoremResponse 由客户端调用, 将服务端返回的结果格式化成LoremResponse对象
func DecodeLoremResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	fmt.Println("ppppppppppppppppp")
	fmt.Printf("%+v\n", resp.Body)
	var response LoremResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// DecodeLoremClientRequest 由客户端调用, 将对客户端的请求转换成LoremRequest对象
// 注意对客户端的请求内容在POST请求体中, 而不是restful路径中.
func DecodeLoremClientRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("0000000000000000000")
	fmt.Printf("%+v\n", r)
	var request LoremRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	fmt.Println("1111111111111111111")
	return request, nil
}
