package transport

import (
	"context"
	"encoding/json"
	"net/http"

	endpoint "github.com/lukyth/try-go-kit/svc-string/pkg/endpoint"
)

// DecodeUppercaseRequest decodes a JSON-encoded uppercase request from the HTTP request body.
func DecodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// DecodeCountRequest decodes a JSON-encoded count request from the HTTP request body.
func DecodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// EncodeResponse encodes the response as JSON to the response writer.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
