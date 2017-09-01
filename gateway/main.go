package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	// r.HandleFunc("/add", addHandler)
	r.HandleFunc("/string", stringHandler)

	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", r))
}

func stringHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/count"

	req, err := http.NewRequest("POST", url, r.Body)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	w.Write([]byte(string(body)))
}
