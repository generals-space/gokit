package usermanager

import (
	"log"
	"net/http"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
	transport_http "github.com/go-kit/kit/transport/http"
)

// StartHTTPTransport 启动http transport
func StartHTTPTransport(srv *UserManager) {
	log.Println("starting user manager http transport...")
	getUserHandler := transport_http.NewServer(
		makeGetUserEndpoint(srv),
		decodeHTTPRequest("GetUser"),
		encodeHTTPResponse,
	)
	addUserHandler := transport_http.NewServer(
		makeGetUserEndpoint(srv),
		decodeHTTPRequest("AddUser"),
		encodeHTTPResponse,
	)
	dispatchHandler := transport_http.NewServer(
		makeGetUserEndpoint(srv),
		decodeHTTPRequest("Dispatch"),
		encodeHTTPResponse,
	)
	http.Handle("/user/query", getUserHandler)
	http.Handle("/user/add", addUserHandler)
	http.Handle("/user/dispatch", dispatchHandler)
	log.Fatal(http.ListenAndServe(common.UserManagerHttpTransportAddr, nil))
}
