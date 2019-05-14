package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/generals-space/gokit/06.gokit-playground-example/21.gokit-lorem-logging"
)

func main() {
	/*
		// 将日志打印到文件(注意, 这个文件只包含接口日志, 不包含项目启动日志等其他信息)
		logfile, err := os.OpenFile("./golorem.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer logfile.Close()
		// 我尝试了下, 下面两句都可以将日志输出到文件, 可能`NewSyncWriter()`可以异步写入吧.
		// target := logfile
		target := log.NewSyncWriter(logfile)
	*/
	// 将日志打印到标准输出
	target := os.Stdout

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(target)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// 创建业务服务
	var svc lorem_logging.Service
	svc = lorem_logging.LoremService{}
	// 主要是这一行, 其他与原来的相同.
	svc = lorem_logging.LoggingMiddleware(logger)(svc)
	// 封装endpoint接口, 一个endpoint其实就是一个接口
	endpoints := lorem_logging.Endpoints{
		LoremEndpoint: lorem_logging.MakeLoremEndpoint(svc),
	}

	ctx := context.Background()
	// 将go-kit类型的endpoint接口转换成http标准库接口
	handler := lorem_logging.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", handler))
}
