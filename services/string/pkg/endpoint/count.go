package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/lukyth/try-go-kit/services/string/pkg/service"
)

// MakeCountEndpoint constructs a Count endpoint wrapping the service.
func MakeCountEndpoint(svc service.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}

// CountRequest collects the request parameters for the Count method.
type CountRequest struct {
	S string `json:"s"`
}

// CountResponse collects the response values for the Count method.
type CountResponse struct {
	V int `json:"v"`
}
