package department

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

// StartHTTPTransport ...
func StartHTTPTransport(srv *DepartmentManager) {
	log.Println("starting department manager http transport...")

	listHandler := transport_http.NewServer(
		makeListEndpoint(srv),
		decodeHTTPRequest(&common.Empty{}),
		encodeHTTPResponse,
	)
	createHandler := transport_http.NewServer(
		makeCreateEndpoint(srv),
		decodeHTTPRequest(&common.Department{}),
		encodeHTTPResponse,
	)
	http.Handle("/department/list", listHandler)
	http.Handle("/department/create", createHandler)
	// log.Fatal(http.ListenAndServe(common.DepartmentHttpTransportAddr, nil))
}
