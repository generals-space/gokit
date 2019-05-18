package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"

	"github.com/generals-space/gokit/06.gokit-playground-example/31.gokit-lorem-ratelimit"
)

func main() {
	// 创建业务服务
	var svc lorem_rate_limit.Service
	svc = lorem_rate_limit.LoremService{}

	// 添加限流中间件, 1s间隔, 桶中5个令牌
	limiter := rate.NewLimiter(1, 5)
	endp := lorem_rate_limit.MakeLoremEndpoint(svc) // 此句保持不变
	endp = ratelimit.NewErroringLimiter(limiter)(endp)

	// 封装endpoint接口, 一个endpoint其实就是一个接口
	endpoints := lorem_rate_limit.Endpoints{
		LoremEndpoint: endp,
	}

	ctx := context.Background()
	// 将go-kit类型的endpoint接口转换成http标准库接口
	handler := lorem_rate_limit.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", handler))
}
