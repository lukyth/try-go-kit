package service

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(StringService) StringService
