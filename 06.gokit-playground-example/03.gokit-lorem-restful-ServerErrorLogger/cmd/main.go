package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/generals-space/gokit/06.gokit-playground-example/03.gokit-lorem-restful-ServerErrorLogger"
)

func main() {
	// 创建业务服务
	var svc lorem_restful.Service
	svc = lorem_restful.LoremService{}
	// 封装endpoint接口, 一个endpoint其实就是一个接口
	endpoints := lorem_restful.Endpoints{
		LoremEndpoint: lorem_restful.MakeLoremEndpoint(svc),
	}

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	ctx := context.Background()
	// 将go-kit类型的endpoint接口转换成http标准库接口
	handler := lorem_restful.MakeHTTPHandler(ctx, endpoints, logger)

	// 提供标准http服务
	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", handler))
}
