package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/juju/ratelimit"

	"gokit/pkg/lorem_rate_limit"
)

func main() {
	// 创建业务服务
	var svc lorem_rate_limit.Service
	svc = lorem_rate_limit.LoremService{}

	// 添加限流中间件, 1s间隔, 桶中3个令牌
	// 注意这里!
	rlbucket := ratelimit.NewBucket(1*time.Second, 3)
	endp := lorem_rate_limit.MakeLoremEndpoint(svc)
	endp = lorem_rate_limit.NewTokenBucketLimiter(rlbucket)(endp)

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
