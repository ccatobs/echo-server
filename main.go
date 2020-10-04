package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	address := ":" + os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(dump))
	})

	log.Fatal(http.ListenAndServe(address, nil))
}
