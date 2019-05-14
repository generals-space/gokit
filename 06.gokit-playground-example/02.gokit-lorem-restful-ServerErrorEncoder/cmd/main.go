package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/generals-space/gokit/06.gokit-playground-example/02.gokit-lorem-restful-ServerErrorEncoder"
)

func main() {
	// 创建业务服务
	var svc lorem_restful.Service
	svc = lorem_restful.LoremService{}
	// 封装endpoint接口, 一个endpoint其实就是一个接口
	endpoints := lorem_restful.Endpoints{
		LoremEndpoint: lorem_restful.MakeLoremEndpoint(svc),
	}

	ctx := context.Background()
	// 将go-kit类型的endpoint接口转换成http标准库接口
	handler := lorem_restful.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server at port 8080")
	fmt.Println(http.ListenAndServe(":8080", handler))
}
