package department

import (
	"log"
	"net/http"

	transport_http "github.com/go-kit/kit/transport/http"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)

// StartHTTPTransport ...
func StartHTTPTransport(srv *DepartmentManager) {
	listHandler := transport_http.NewServer(
		makeListEndpoint(srv),
		decodeHTTPRequest("List"),
		encodeHTTPResponse,
	)
	createHandler := transport_http.NewServer(
		makeCreateEndpoint(srv),
		decodeHTTPRequest("Create"),
		encodeHTTPResponse,
	)

	http.Handle("/department/list", listHandler)
	http.Handle("/department/create", createHandler)
	log.Fatal(http.ListenAndServe(common.DepartmentHttpTransportAddr, nil))
}
