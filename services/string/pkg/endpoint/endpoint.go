package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/lukyth/try-go-kit/services/string/pkg/service"
)

// Set collects all of the endpoints that compose an add service. It's meant to
// be used as a helper struct, to collect all of the endpoints into a single
// parameter.
type Set struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
}

// New returns a Set that wraps the provided server, and wires in all of the
// expected endpoint middlewares via the various parameters.
func New(svc service.StringService) Set {
	return Set{
		UppercaseEndpoint: MakeUppercaseEndpoint(svc),
		CountEndpoint:     MakeCountEndpoint(svc),
	}
}
