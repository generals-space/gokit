package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	"gokit/pkg/lorem_tracing"
)

func main() {
	var (
		serverAddr = os.Getenv("SERVER_ADDR")
		serverPort = os.Getenv("SERVER_PORT")
		zipkinURL  = os.Getenv("ZIPKIN_URL")
	)

	reporter := zipkinhttp.NewReporter(zipkinURL)
	defer reporter.Close()
	zipkinTracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var svc lorem_tracing.Service
	svc = lorem_tracing.LoremService{}

	loremEndpoint := lorem_tracing.MakeLoremEndpoint(svc)
	loremEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "lorem-endpoint")(loremEndpoint)
	endpoints := lorem_tracing.Endpoints{
		LoremEndpoint: loremEndpoint,
	}

	ctx := context.Background()
	handler := lorem_tracing.MakeHTTPHandler(ctx, endpoints, zipkinTracer)

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(serverAddr+":"+serverPort, handler))
}
