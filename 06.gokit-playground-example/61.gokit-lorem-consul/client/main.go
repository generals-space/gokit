package main

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"

	"github.com/generals-space/gokit/06.gokit-playground-example/61.gokit-lorem-consul"
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
		lorem_consul.EncodeLoremRequest,
		lorem_consul.DecodeLoremResponse,
	).Endpoint(), nil, nil
}

func main() {
	consulAddr := "consul-svc"
	consulPort := "8500"

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
	endpointer := sd.NewEndpointer(instancer, LoremFactory, logger)

	// 这里是我自己加的.
	// 如果没有在consul中发现目标服务, 则return, 不再尝试发送请求.
	// 不过如果服务不存在, 应该到不了这里, 在instancer处就会报错的? 待验证<???>.
	endpointList, err := endpointer.Endpoints()
	if len(endpointList) == 0 {
		fmt.Println("endpoints not found")
		return
	}

	// 负载均衡器
	balancer := lb.NewRoundRobin(endpointer)
	loremEndpoint := lb.Retry(1, time.Millisecond*500, balancer)

	loremRequest := lorem_consul.LoremRequest{
		RequestType: "Sentence",
		Min:         5,
		Max:         20,
	}
	// 通过endpoint端点对象发送请求, 超时时间为5s
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 原作代码中是将得到的endpoint对象当作了一个路由处理函数, 挂载到了http server上.
	// 但这样的话, 客户端就相当于是一个路由转发的工具, 而不是具体的某一个服务了.
	// 这里我们不那么做, 而是直接通过endpoint对象来执行一些操作.
	msg, err := loremEndpoint(ctx, loremRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(msg)
}
