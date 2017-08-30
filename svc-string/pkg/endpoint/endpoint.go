package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lukyth/try-go-kit/svc-string/pkg/service"
)

// MakeUppercaseEndpoint constructs an Uppercase endpoint wrapping the service.
func MakeUppercaseEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}
		return UppercaseResponse{v, ""}, nil
	}
}

// MakeCountEndpoint constructs a Count endpoint wrapping the service.
func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}

// UppercaseRequest collects the request parameters for the Uppercase method.
type UppercaseRequest struct {
	S string `json:"s"`
}

// UppercaseResponse collects the response values for the Uppercase method.
type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

// CountRequest collects the request parameters for the Count method.
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse collects the response values for the Count method.
type CountResponse struct {
	V int `json:"v"`
}
