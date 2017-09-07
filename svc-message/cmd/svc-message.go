package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

type httpHandler struct{}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Service Message"))
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", httpHandler{}))
}
