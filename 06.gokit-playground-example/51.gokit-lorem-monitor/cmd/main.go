package main

import (
	"context"
	"fmt"
	"net/http"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/generals-space/gokit/06.gokit-playground-example/51.gokit-lorem-monitor"
)

func main() {
	//declare metrics
	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "ru_rocker",
		Subsystem: "lorem_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "ru_rocker",
		Subsystem: "lorem_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// 创建业务服务
	var svc lorem_metrics.Service
	svc = lorem_metrics.LoremService{}
	svc = lorem_metrics.MetricsMiddleware(requestCount, requestLatency)(svc)

	// 封装endpoint接口, 一个endpoint其实就是一个接口
	endpoints := lorem_metrics.Endpoints{
		LoremEndpoint: lorem_metrics.MakeLoremEndpoint(svc),
	}

	ctx := context.Background()
	// 将go-kit类型的endpoint接口转换成http标准库接口
	handler := lorem_metrics.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", handler))
}
