package department

import(
	"context"
	"net/http"
	"encoding/json"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)
// type gokitDecodeFunc func(context.Context, *http.Request)(interface{}, error)
func decodeHTTPRequest(reqType string) func(context.Context, *http.Request) (interface{}, error) {
	var request interface{}
	switch reqType {
	case "List":
		request = &common.Empty{}
	case "Create":
		request = &common.Department{}
	}

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
