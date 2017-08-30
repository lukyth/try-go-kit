package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/lukyth/try-go-kit/svc-string/pkg/endpoint"
	"github.com/lukyth/try-go-kit/svc-string/pkg/service"
	"github.com/lukyth/try-go-kit/svc-string/pkg/transport"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.New(logger)
	endpoints := endpoint.New(svc)
	httpHandler := transport.NewHTTPHandler(endpoints)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", httpHandler))
}
