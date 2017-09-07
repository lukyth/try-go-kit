package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lukyth/try-go-kit/services/message/pkg/endpoints"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()
	m.Handle("/get_messages", httptransport.NewServer(
		endpoints.GetMessagesEndpoint,
		DecodeGetMessagesRequest,
		EncodeGetMessagesResponse,
	))
	m.Handle("/get_message", httptransport.NewServer(
		endpoints.GetMessageEndpoint,
		DecodeGetMessageRequest,
		EncodeGetMessageResponse,
	))
	m.Handle("/post_message", httptransport.NewServer(
		endpoints.PostMessageEndpoint,
		DecodePostMessageRequest,
		EncodePostMessageResponse,
	))
	return m
}

// DecodeGetMessagesRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeGetMessagesRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.GetMessagesRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeGetMessagesResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeGetMessagesResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodeGetMessageRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodeGetMessageRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.GetMessageRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodeGetMessageResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodeGetMessageResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}

// DecodePostMessageRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body. Primarily useful in a server.
func DecodePostMessageRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	req = endpoints.PostMessageRequest{}
	err = json.NewDecoder(r.Body).Decode(&r)
	return req, err
}

// EncodePostMessageResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer. Primarily useful in a server.
func EncodePostMessageResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return err
}
