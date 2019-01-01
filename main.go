package main

import (
	_ "chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
	"net/http"
)

func main() {

	mux := httprouter.New()
	mux.GET("/", Index)
	mux.GET("/threads/:id", FindThread)
	mux.GET("/signup", Signup)
	mux.POST("/signup_account", SignupAccount)
	//files := http.FileServer(http.Dir("/public"))
	//mux.Handle("/static/", http.StripPrefix("/static/", files))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	http2.ConfigureServer(server, &http2.Server{})
	server.ListenAndServe()
}
