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
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	"github.com/generals-space/gokit/07.gokit链路追踪/02.gokit-lorem-tracing-consul"
)

func connectConsul(addr, port string) (client sdconsul.Client, err error) {
	consulConfig := api.DefaultConfig()

	consulConfig.Address = "http://" + addr + ":" + port
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		return
	}
	client = sdconsul.NewClient(consulClient)
	return
}

func buildFactory(zipkinTracer *zipkin.Tracer) sd.Factory {
	// LoremFactory endpoint端点的动态构造工厂.
	// 由sd.NewEndpointer()调用, 调用时传入从consul服务实例管理器中选出的服务地址`instance`,
	// 服务地址一般是IP:Port, 需要此函数根据服务地址创建完成的端点请求地址.
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}

		tgt, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		tgt.Path = "/lorem"

		zipkinClientTrace := kitzipkin.HTTPClientTrace(zipkinTracer)
		options := []httptransport.ClientOption{
			zipkinClientTrace,
		}

		endp := httptransport.NewClient(
			"POST",
			tgt,
			lorem_tracing.EncodeLoremRequest,
			lorem_tracing.DecodeLoremResponse,
			options...,
		).Endpoint()
		endp = kitzipkin.TraceEndpoint(zipkinTracer, "http-client")(endp)
		return endp, nil, nil
	}
}

func main() {
	var (
		consulAddr    = "consul-svc"
		consulPort    = "8500"
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
		zipkinURL     = os.Getenv("ZIPKIN_URL")
	)

	reporter := zipkinhttp.NewReporter(zipkinURL)
	defer reporter.Close()
	zipkinTracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	client, err := connectConsul(consulAddr, consulPort)
	if err != nil {
		panic(err)
	}

	serviceName := "lorem"
	tags := []string{"lorem", "ru-rocker"}
	// NewInstancer返回一个服务实例管理器, 包含指定条件(名称, 标签都符合)的服务实例的地址.
	instancer := sdconsul.NewInstancer(client, logger, serviceName, tags, true)

	// NewEndpointer返回一个端点管理器.
	// 此管理器监听instancer内服务实例的的变化(如掉线, 新增服务实例等), 通过Factory动态更新创建的endpointer.
	endpointer := sd.NewEndpointer(instancer, buildFactory(zipkinTracer), logger)

	// 负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	loremEndpoint := lb.Retry(1, time.Millisecond*500, balancer)

	// POST /sd-lorem
	// Payload: {"requestType":"word", "min":10, "max":10}
	r := mux.NewRouter()
	r.Methods("POST").Path("/sd-lorem").Handler(httptransport.NewServer(
		loremEndpoint,
		lorem_tracing.DecodeLoremClientRequest,
		lorem_tracing.EncodeResponse,
	))
	loremEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "http-client")(loremEndpoint)

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, r))
}
