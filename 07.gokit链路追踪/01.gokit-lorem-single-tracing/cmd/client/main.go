package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"

	"gokit/pkg/lorem_tracing"
)

func main() {
	var (
		advertiseAddr = os.Getenv("SERVER_ADDR")
		advertisePort = os.Getenv("SERVER_PORT")
		zipkinURL     = os.Getenv("ZIPKIN_URL")
	)

	reporter := zipkinhttp.NewReporter(zipkinURL)
	defer reporter.Close()
	zipkinTracer, err := zipkin.NewTracer(reporter)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	baseURL := "http://server:8080"
	urlObj, _ := url.Parse(baseURL)
	urlObj.Path = "/lorem"

	zipkinClientTrace := kitzipkin.HTTPClientTrace(zipkinTracer)
	options := []httptransport.ClientOption{
		zipkinClientTrace,
	}

	loremEndpoint := httptransport.NewClient(
		"POST",
		urlObj,
		lorem_tracing.EncodeLoremRequest,
		lorem_tracing.DecodeLoremResponse,
		options...,
	).Endpoint()
	loremEndpoint = kitzipkin.TraceEndpoint(zipkinTracer, "http-client")(loremEndpoint)

	// POST /sd-lorem
	// Payload: {"requestType":"word", "min":10, "max":10}
	r := mux.NewRouter()
	r.Methods("POST").Path("/sd-lorem").Handler(httptransport.NewServer(
		loremEndpoint,
		lorem_tracing.DecodeLoremClientRequest,
		lorem_tracing.EncodeResponse,
	))

	// 提供标准http服务
	fmt.Println("Starting server")
	fmt.Println(http.ListenAndServe(advertiseAddr+":"+advertisePort, r))
}
