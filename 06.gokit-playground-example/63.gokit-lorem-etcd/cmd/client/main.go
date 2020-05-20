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
	sdetcdv3 "github.com/go-kit/kit/sd/etcdv3"

	"gokit/pkg/lorem_etcd"
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
		lorem_etcd.EncodeLoremRequest,
		lorem_etcd.DecodeLoremResponse,
	).Endpoint(), nil, nil
}

func main() {
	var (
		etcdURL       = os.Getenv("ETCD_URL")
		etcdPrefix    = os.Getenv("ETCD_PREFIX")
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
	)

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	client, err := lorem_etcd.ConnectEtcd(etcdURL)
	if err != nil {
		panic(err)
	}

	// NewInstancer返回一个服务实例管理器, 包含指定条件(名称, 标签都符合)的服务实例的地址.
	instancer, err := sdetcdv3.NewInstancer(client, etcdPrefix, logger)
	if err != nil {
		panic(err)
	}
	
	// NewEndpointer返回一个端点管理器.
	// 此管理器监听instancer内服务实例的的变化(如掉线, 新增服务实例等), 通过Factory动态更新创建的endpointer.
	endpointer := sd.NewEndpointer(instancer, LoremFactory, logger)

	// 这里是我自己加的. 如果没有在consul中发现目标服务, 则return, 不再尝试发送请求.
	// 其实在实际场景中不应加这句, 因为如果使用compose同时启动server与client时,
	// server可能还未来得及启动, client运行到这里一定会退出.
	// endpointList, err := endpointer.Endpoints()
	// if len(endpointList) == 0 {
	// 	fmt.Println("endpoints not found")
	// 	return
	// }

	// 负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	loremEndpoint := lb.Retry(1, time.Millisecond*500, balancer)

	// POST /sd-lorem
	// Payload: {"requestType":"word", "min":10, "max":10}
	r := mux.NewRouter()
	r.Methods("POST").Path("/sd-lorem").Handler(httptransport.NewServer(
		loremEndpoint,
		lorem_etcd.DecodeLoremClientRequest,
		lorem_etcd.EncodeResponse,
	))

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, r))
}
