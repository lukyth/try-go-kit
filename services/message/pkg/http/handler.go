package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/lukyth/try-go-kit/services/message/pkg/endpoints"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/messages").Handler(httptransport.NewServer(
		endpoints.GetMessagesEndpoint,
		DecodeGetMessagesRequest,
		EncodeResponse,
	))
	r.Methods("GET").Path("/messages/{id}").Handler(httptransport.NewServer(
		endpoints.GetMessageEndpoint,
		DecodeGetMessageRequest,
		EncodeResponse,
	))
	r.Methods("POST").Path("/messages").Handler(httptransport.NewServer(
		endpoints.PostMessageEndpoint,
		DecodePostMessageRequest,
		EncodeResponse,
	))
	return r
}

// DecodeGetMessagesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeGetMessagesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return endpoints.GetMessagesRequest{}, nil
}

// DecodeGetMessageRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeGetMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errors.New("inconsistent mapping between route and handler (programmer error)")
	}
	return endpoints.GetMessageRequest{MID: id}, nil
}

// DecodePostMessageRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodePostMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.PostMessageRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// EncodeResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(response)
	return err
}
