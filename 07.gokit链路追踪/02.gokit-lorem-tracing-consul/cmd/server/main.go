package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	"gokit/pkg/lorem_tracing"
)

func main() {
	var (
		// 由于consul服务运行在docker或compose, 所以这两个地址一定要正确.
		consulAddr    = "consul-svc"
		consulPort    = "8500"
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
		serviceName   = "lorem"
		zipkinURL     = os.Getenv("ZIPKIN_URL")
	)

	reporter := zipkinhttp.NewReporter(zipkinURL)
	defer reporter.Close()
	zipkinTracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var svc lorem_tracing.Service
	svc = lorem_tracing.LoremService{}

	loremEndpoint := lorem_tracing.MakeLoremEndpoint(svc)
	loremEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "lorem-endpoint")(loremEndpoint)
	healthEndpoint := lorem_tracing.MakeHealthEndpoint(svc)
	healthEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "health-endpoint")(healthEndpoint)
	endpoints := lorem_tracing.Endpoints{
		LoremEndpoint:  loremEndpoint,
		HealthEndpoint: healthEndpoint,
	}

	// 注册服务
	registrar := lorem_tracing.Register(consulAddr, consulPort, advertiseAddr, advertisePort, serviceName)
	// 将go-kit类型的endpoint接口转换成http标准库接口
	registrar.Register()

	ctx := context.Background()
	handler := lorem_tracing.MakeHTTPHandler(ctx, endpoints, zipkinTracer)

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, handler))

	registrar.Deregister()
}
