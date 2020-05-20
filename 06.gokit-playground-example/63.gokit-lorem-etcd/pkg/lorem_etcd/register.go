package lorem_etcd

import (
	"context"
	"os"
	"time"

	gkLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	sdetcdv3 "github.com/go-kit/kit/sd/etcdv3"
)

func ConnectEtcd(etcdURL string) (client sdetcdv3.Client, err error) {
	//etcd的连接参数
	options := sdetcdv3.ClientOptions{
		DialTimeout:   time.Second * 3,
		DialKeepAlive: time.Second * 3,
	}
	ctx := context.Background()
	//创建etcd连接
	client, err = sdetcdv3.NewClient(ctx, []string{etcdURL}, options)
	return
}

// Register 创建sd.Registrar注册中心对象, 该对象拥有`Register()`和`Deregister()`方法, 可以实现注册与注销功能.
func Register(client sdetcdv3.Client, key, value string) (registrar sd.Registrar) {
	var logger gkLog.Logger
	{
		logger = gkLog.NewLogfmtLogger(os.Stderr)
		logger = gkLog.With(logger, "ts", gkLog.DefaultTimestampUTC)
		logger = gkLog.With(logger, "caller", gkLog.DefaultCaller)
	}

	registrar = sdetcdv3.NewRegistrar(client, sdetcdv3.Service{
		Key:   key,
		Value: value,
		// 注册服务的健康检查机制关键点在这里, go-kit会为注册操作设置一个默认值(3, 10)
		TTL: sdetcdv3.NewTTLOption(time.Second*5, time.Second*10),
	}, logger)
	return
}
