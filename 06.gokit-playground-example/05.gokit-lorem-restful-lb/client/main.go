package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/generals-space/gokit/06.gokit-playground-example/05.gokit-lorem-restful-lb"
)

// LoremFactory endpoint端点的动态构造工厂.
// 由sd.NewEndpointer()调用, 调用时传入从consul服务实例管理器中选出的服务地址`instance`,
// 服务地址一般是IP:Port, 需要此函数根据服务地址创建完成的端点请求地址.
func LoremFactory(instance string) (endpoint.Endpoint, io.Closer, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	tgt, err := url.Parse(instance)
	if err != nil {
		return nil, nil, err
	}
	tgt.Path = "/lorem"

	return httptransport.NewClient(
		"POST",
		tgt,
		lorem_restful.EncodeLoremRequest,
		lorem_restful.DecodeLoremResponse,
	).Endpoint(), nil, nil
}

func makeLoremClientEndpoint(instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	tgt, err := url.Parse(instance)
	if err != nil {
		return nil
	}
	tgt.Path = "/lorem"

	return httptransport.NewClient(
		"POST",
		tgt,
		lorem_restful.EncodeLoremRequest,
		lorem_restful.DecodeLoremResponse,
	).Endpoint()
}

func main() {
	var (
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
		serverAddrs   = os.Getenv("SERVER_LIST")
	)

	// 本例重点: 手动构建endpointer的方法, 有两种.
	// 第一种, 通过instancer对象和sd.Factory对象得到endpointer
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	serverList := strings.Split(serverAddrs, ",")
	instancer := sd.FixedInstancer(serverList)
	endpointer := sd.NewEndpointer(instancer, LoremFactory, logger)

	/*
		// 第二种, 直接创建Endpointer对象
		endpointer := sd.FixedEndpointer{}
		for _, serverAddr := range strings.Split(serverAddrs, ",") {
			endp := makeLoremClientEndpoint(serverAddr)
			endpointer = append(endpointer, endp)
		}
	*/
	// 负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	loremEndpoint := lb.Retry(1, time.Millisecond*500, balancer)

	// POST /sd-lorem
	// Payload: {"requestType":"word", "min":10, "max":10}
	r := mux.NewRouter()
	r.Methods("POST").Path("/sd-lorem").Handler(httptransport.NewServer(
		loremEndpoint,
		lorem_restful.DecodeLoremClientRequest,
		lorem_restful.EncodeResponse,
	))

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, r))
}
