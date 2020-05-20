package usermanager

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	transport_http "github.com/go-kit/kit/transport/http"

	"gokit/common"
)

func decodeHTTPRequest(request interface{}) func(context.Context, *http.Request) (interface{}, error) {
	return func(_ context.Context, r *http.Request) (interface{}, error) {
		if err := json.NewDecoder(r.Body).Decode(request); err != nil {
			return nil, err
		}
		return request, nil
	}
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// StartHTTPTransport 启动http transport
func StartHTTPTransport(srv *UserManager) {
	log.Println("starting user manager http transport...")

	listHandler := transport_http.NewServer(
		makeListEndpoint(srv),
		decodeHTTPRequest(&common.Empty{}),
		encodeHTTPResponse,
	)
	dispatchHandler := transport_http.NewServer(
		makeDispatchEndpoint(srv),
		decodeHTTPRequest(&common.DispatchRequest{}),
		encodeHTTPResponse,
	)
	http.Handle("/user/list", listHandler)
	http.Handle("/user/dispatch", dispatchHandler)
	// log.Fatal(http.ListenAndServe(common.UserManagerHttpTransportAddr, nil))
}
