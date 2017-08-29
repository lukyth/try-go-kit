package middleware

import "github.com/lukyth/try-go-kit/svc-string/pkg/service"

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(service.StringService) service.StringService
