package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/lukyth/try-go-kit/services/message/pkg/endpoints"
	transport "github.com/lukyth/try-go-kit/services/message/pkg/http"
	"github.com/lukyth/try-go-kit/services/message/pkg/service"
)

type httpHandler struct{}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.New()
	ep := endpoints.New(svc)
	httpHandler := transport.NewHTTPHandler(ep)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", httpHandler))
}
