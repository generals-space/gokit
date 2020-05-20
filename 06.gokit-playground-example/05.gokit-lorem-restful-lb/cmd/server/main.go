package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"gokit/pkg/lorem_restful"
)

var (
	cmdFlags      = flag.NewFlagSet("server", flag.ExitOnError)
	advertiseAddr = "localhost"
	advertisePort = "8080"
)

func init() {
	cmdFlags.StringVar(&advertiseAddr, "addr", "localhost", "监听地址")
	cmdFlags.StringVar(&advertisePort, "port", "8080", "监听端口")
	cmdFlags.Parse(os.Args[1:])
}

func main() {
	var svc lorem_restful.Service
	svc = lorem_restful.LoremService{}

	loremEndpoint := lorem_restful.MakeLoremEndpoint(svc)
	endpoints := lorem_restful.Endpoints{
		LoremEndpoint: loremEndpoint,
	}

	ctx := context.Background()
	handler := lorem_restful.MakeHTTPHandler(ctx, endpoints)

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, handler))
}
