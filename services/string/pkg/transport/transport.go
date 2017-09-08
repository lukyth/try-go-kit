package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/lukyth/try-go-kit/services/string/pkg/endpoint"
)

// NewHTTPHandler returns an HTTP handler that makes a set of endpoints
// available on predefined paths.
func NewHTTPHandler(endpoints endpoint.Set) http.Handler {
	m := http.NewServeMux()
	m.Handle("/uppercase", httptransport.NewServer(
		endpoints.UppercaseEndpoint,
		DecodeUppercaseRequest,
		EncodeResponse,
	))
	m.Handle("/count", httptransport.NewServer(
		endpoints.CountEndpoint,
		DecodeCountRequest,
		EncodeResponse,
	))
	m.Handle("/metrics", promhttp.Handler())
	return m
}

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
