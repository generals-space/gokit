package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/generals-space/gokit/06.gokit-playground-example/04.gokit-lorem-restful-client"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	baseURL := "http://server:8080"
	urlObj, err := url.Parse(baseURL)
	router := mux.NewRouter()
	routes := lorem_restful.Initialize(router)

	loremEndpoint := httptransport.NewClient(
		"POST", // 因为在EncodeRequest已经有method赋值操作, 所以其实这里可以为空.
		urlObj,
		lorem_restful.EncodeLoremRequest(routes.Lorem),
		lorem_restful.DecodeLoremResponse,
	).Endpoint()

	endpoints := lorem_restful.Endpoints{
		LoremEndpoint: loremEndpoint,
	}

	ctx := context.Background()
	msg, err := endpoints.Lorem(ctx, "Sentence", 5, 20)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(msg)
}
