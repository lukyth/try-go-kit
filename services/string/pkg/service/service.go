package service

import (
	"errors"
	"strings"

	"github.com/go-kit/kit/log"
)

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger log.Logger) StringService {
	var svc StringService
	{
		svc = NewStringService()
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware()(svc)
	}
	return svc
}

// NewStringService returns a naïve, stateless implementation of StringService.
func NewStringService() StringService {
	return stringService{}
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")
