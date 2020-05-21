package lorem_tracing

import (
	"context"
	"errors"
	"net/http"

	kitzipkin "github.com/go-kit/kit/tracing/zipkin"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	gozipkin "github.com/openzipkin/zipkin-go"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler ...
func MakeHTTPHandler(_ context.Context, endpoint Endpoints, zipkinTracer *gozipkin.Tracer) http.Handler {
	r := mux.NewRouter()

	zipkinServerTrace := kitzipkin.HTTPServerTrace(zipkinTracer, kitzipkin.Name("http-transport"))

	options := []httptransport.ServerOption{
		zipkinServerTrace,
	}

	//POST /lorem/{type}/{min}/{max}
	r.Methods("POST").Path("/lorem/{type}/{min}/{max}").Handler(httptransport.NewServer(
		endpoint.LoremEndpoint,
		DecodeLoremRequest,
		EncodeResponse,
		options...,
	))

	return r
}
