package main

import (
	"main/buz"
	"net/http"
)

func main() {
	mux :=http.NewServeMux()
	mux.HandleFunc("/user/authenticate", buz.authenticate)

	server := &http.Server{
		Addr:	"127.0.0.1:80",
		Handler: mux,
	}
	server.ListenAndServe()
}
