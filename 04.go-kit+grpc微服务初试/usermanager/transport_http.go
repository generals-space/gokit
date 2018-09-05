package usermanager

import (
	"log"
	"net/http"
	"context"
	"encoding/json"
	
	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	transport_http "github.com/go-kit/kit/transport/http"
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
	getUserHandler := transport_http.NewServer(
		makeGetUserEndpoint(srv),
		decodeHTTPRequest(&common.GetUserRequest{}),
		encodeHTTPResponse,
	)
	addUserHandler := transport_http.NewServer(
		makeAddUserEndpoint(srv),
		decodeHTTPRequest(&common.AddUserRequest{}),
		encodeHTTPResponse,
	)
	dispatchHandler := transport_http.NewServer(
		makeDispatchEndpoint(srv),
		decodeHTTPRequest(&common.DispatchRequest{}),
		encodeHTTPResponse,
	)
	http.Handle("/user/query", getUserHandler)
	http.Handle("/user/add", addUserHandler)
	http.Handle("/user/dispatch", dispatchHandler)
	log.Fatal(http.ListenAndServe(common.UserManagerHttpTransportAddr, nil))
}
