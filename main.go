package main

import (
	_ "chit-chat/data"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	//files := http.FileServer(http.Dir("/public"))
	//mux.Handle("/static/", http.StripPrefix("/static/", files))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	mux.HandleFunc("/", index)
	server.ListenAndServe()
}
