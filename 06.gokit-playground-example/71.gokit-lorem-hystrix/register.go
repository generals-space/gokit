package lorem_hystrix

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	gkLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	sdconsul "github.com/go-kit/kit/sd/consul"
	"github.com/hashicorp/consul/api"
)

// Register 创建sd.Registrar注册中心对象, 该对象拥有`Register()`和`Deregister()`方法, 可以实现注册与注销功能.
func Register(consulAddress, consulPort, advertiseAddress, advertisePort string) (registar sd.Registrar) {
	var logger gkLog.Logger
	{
		logger = gkLog.NewLogfmtLogger(os.Stderr)
		logger = gkLog.With(logger, "ts", gkLog.DefaultTimestampUTC)
		logger = gkLog.With(logger, "caller", gkLog.DefaultCaller)
	}

	var client sdconsul.Client
	{
		consulConfig := api.DefaultConfig()
		consulConfig.Address = consulAddress + ":" + consulPort
		consulClient, err := api.NewClient(consulConfig)
		if err != nil {
			logger.Log("err", err)
			os.Exit(1)
		}
		client = sdconsul.NewClient(consulClient)
	}

	check := api.AgentServiceCheck{
		HTTP:     "http://" + advertiseAddress + ":" + advertisePort + "/health",
		Interval: "10s",
		Timeout:  "1s",
		Notes:    "Basic health checks",
	}

	port, _ := strconv.Atoi(advertisePort)
	// 确保服务ID唯一
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Intn(100)
	asr := api.AgentServiceRegistration{
		ID:      "lorem" + strconv.Itoa(num),
		Name:    "lorem",
		Address: advertiseAddress,
		Port:    port,
		Tags:    []string{"lorem", "ru-rocker"},
		Check:   &check,
	}
	registar = sdconsul.NewRegistrar(client, &asr, logger)
	return
}
