package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/lukyth/try-go-kit/services/message/pkg/service"
)

// Endpoints collects all of the endpoints that compose a message service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
type Endpoints struct {
	GetMessagesEndpoint endpoint.Endpoint
	GetMessageEndpoint  endpoint.Endpoint
	PostMessageEndpoint endpoint.Endpoint
}

// GetMessagesRequest collects the request parameters for the GetMessages method.
type GetMessagesRequest struct{}

// GetMessagesResponse collects the response values for the GetMessages method.
type GetMessagesResponse struct {
	M0 []service.Message
	E1 error
}

// GetMessageRequest collects the request parameters for the GetMessage method.
type GetMessageRequest struct {
	MID string
}

// GetMessageResponse collects the response values for the GetMessage method.
type GetMessageResponse struct {
	M0 service.Message
	E1 error
}

// PostMessageRequest collects the request parameters for the PostMessage method.
type PostMessageRequest struct {
	M service.Message
}

// PostMessageResponse collects the response values for the PostMessage method.
type PostMessageResponse struct {
	E0 error
}

// New return all endpoints.
func New(svc service.MessageService) (ep Endpoints) {
	ep.GetMessagesEndpoint = MakeGetMessagesEndpoint(svc)
	ep.GetMessageEndpoint = MakeGetMessageEndpoint(svc)
	ep.PostMessageEndpoint = MakePostMessageEndpoint(svc)
	return ep
}

// MakeGetMessagesEndpoint returns an endpoint that invokes GetMessages on the service.
// Primarily useful in a server.
func MakeGetMessagesEndpoint(svc service.MessageService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, _ interface{}) (interface{}, error) {
		M0, e1 := svc.GetMessages(ctx)
		return GetMessagesResponse{M0: M0, E1: e1}, nil
	}
}

// MakeGetMessageEndpoint returns an endpoint that invokes GetMessage on the service.
// Primarily useful in a server.
func MakeGetMessageEndpoint(svc service.MessageService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMessageRequest)
		M0, e1 := svc.GetMessage(ctx, req.MID)
		return GetMessageResponse{M0: M0, E1: e1}, nil
	}
}

// MakePostMessageEndpoint returns an endpoint that invokes PostMessage on the service.
// Primarily useful in a server.
func MakePostMessageEndpoint(svc service.MessageService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PostMessageRequest)
		e0 := svc.PostMessage(ctx, req.M)
		return PostMessageResponse{E0: e0}, nil
	}
}
