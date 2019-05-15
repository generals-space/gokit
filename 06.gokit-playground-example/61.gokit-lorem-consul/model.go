package lorem_consul

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// decode url path variables into request
func DecodeLoremRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrBadRouting
	}

	vmin, ok := vars["min"]
	if !ok {
		return nil, ErrBadRouting
	}

	vmax, ok := vars["max"]
	if !ok {
		return nil, ErrBadRouting
	}

	min, _ := strconv.Atoi(vmin)
	max, _ := strconv.Atoi(vmax)

	request := LoremRequest{
		RequestType: requestType,
		Min:         min,
		Max:         max,
	}
	return request, nil
}

// decode health check
func decodeHealthRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return HealthRequest{}, nil
}

// encodeResponse is the common method to encode all response types to the client.
func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// Encode request form LoremRequest into existing Lorem Service
// The encode will translate the LoremRequest into /lorem/{requestType}/{min}/{max}
func EncodeLoremRequest(_ context.Context, req *http.Request, request interface{}) error {
	lr := request.(LoremRequest)
	p := "/" + lr.RequestType + "/" + strconv.Itoa(lr.Min) + "/" + strconv.Itoa(lr.Max)
	req.URL.Path += p
	return nil
}

// decode response from Lorem Service
func DecodeLoremResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response LoremResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
