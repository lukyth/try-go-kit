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
	r.HandleFunc("/{serviceName}/{func}", serviceHandler)

	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", r))
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := "http://" + vars["serviceName"] + ":8080/" + vars["func"]

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
