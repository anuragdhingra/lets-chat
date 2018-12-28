package main

import (
	_ "chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"net/http"

)

func main() {

	mux := httprouter.New()
	mux.GET("/", Index)
	mux.GET("/threads/:id", FindThread)
	//files := http.FileServer(http.Dir("/public"))
	//mux.Handle("/static/", http.StripPrefix("/static/", files))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
