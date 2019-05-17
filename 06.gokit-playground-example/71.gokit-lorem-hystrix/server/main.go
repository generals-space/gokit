package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/generals-space/gokit/06.gokit-playground-example/71.gokit-lorem-hystrix"
)

func main() {
	var (
		// 由于consul服务运行在docker或compose, 所以这两个地址一定要正确.
		consulAddr    = "consul-svc"
		consulPort    = "8500"
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
	)

	var svc lorem_hystrix.Service
	svc = lorem_hystrix.LoremService{}

	loremEndpoint := lorem_hystrix.MakeLoremEndpoint(svc)
	healthEndpoint := lorem_hystrix.MakeHealthEndpoint(svc)
	endpoints := lorem_hystrix.Endpoints{
		LoremEndpoint:  loremEndpoint,
		HealthEndpoint: healthEndpoint,
	}

	// 注册服务
	registar := lorem_hystrix.Register(consulAddr, consulPort, advertiseAddr, advertisePort)
	// 将go-kit类型的endpoint接口转换成http标准库接口
	registar.Register()

	ctx := context.Background()
	handler := lorem_hystrix.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, handler))

	registar.Deregister()
}
