package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/lukyth/try-go-kit/services/string/pkg/endpoint"
	"github.com/lukyth/try-go-kit/services/string/pkg/service"
	"github.com/lukyth/try-go-kit/services/string/pkg/transport"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.New(logger)
	endpoints := endpoint.New(svc)
	httpHandler := transport.NewHTTPHandler(endpoints)

	logger.Log("msg", "HTTP", "addr", ":8081")
	logger.Log("err", http.ListenAndServe(":8081", httpHandler))
}
