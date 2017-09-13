package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/lukyth/try-go-kit/services/message/pkg/endpoints"
	"github.com/lukyth/try-go-kit/services/message/pkg/service"
)

// NewHTTPHandler returns a handler that makes a set of endpoints available on
// predefined paths.
func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/messages").Handler(httptransport.NewServer(
		endpoints.GetMessagesEndpoint,
		DecodeGetMessagesRequest,
		encodeResponse,
	))
	r.Methods("GET").Path("/messages/{id}").Handler(httptransport.NewServer(
		endpoints.GetMessageEndpoint,
		DecodeGetMessageRequest,
		encodeResponse,
	))
	r.Methods("POST").Path("/messages").Handler(httptransport.NewServer(
		endpoints.PostMessageEndpoint,
		DecodePostMessageRequest,
		encodeResponse,
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

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// TODO: Make this line work
	e, ok := response.(errorer)
	if ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrNotFound:
		return http.StatusNotFound
	case service.ErrAlreadyExists, service.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
