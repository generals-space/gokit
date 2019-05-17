package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/generals-space/gokit/06.gokit-playground-example/05.gokit-lorem-restful-lb"
)

func main() {
	var (
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
	)

	var svc lorem_consul.Service
	svc = lorem_consul.LoremService{}

	loremEndpoint := lorem_consul.MakeLoremEndpoint(svc)
	endpoints := lorem_consul.Endpoints{
		LoremEndpoint: loremEndpoint,
	}

	ctx := context.Background()
	handler := lorem_consul.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, handler))
}
