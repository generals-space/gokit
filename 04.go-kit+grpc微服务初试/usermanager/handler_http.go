package usermanager

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/generals-space/gokit/04.go-kit+grpc微服务初试/common"
)

/*
func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request common.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeAddUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request common.AddUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeDispatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request common.DispatchRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
*/

// type gokitDecodeFunc func(context.Context, *http.Request)(interface{}, error)
func decodeHTTPRequest(reqType string) func(context.Context, *http.Request) (interface{}, error) {
	var request interface{}
	switch reqType {
	case "GetUser":
		request = &common.GetUserRequest{}
	case "AddUser":
		request = &common.AddUserRequest{}
	case "Dispatch":
		request = &common.DispatchRequest{}
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
