package lorem_etcd

import (
	"context"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// Make Http Handler
func MakeHTTPHandler(_ context.Context, endpoint Endpoints) http.Handler {
	r := mux.NewRouter()

	//POST /lorem/{type}/{min}/{max}
	r.Methods("POST").Path("/lorem/{type}/{min}/{max}").Handler(httptransport.NewServer(
		endpoint.LoremEndpoint,
		DecodeLoremRequest,
		EncodeResponse,
	))

	return r
}
