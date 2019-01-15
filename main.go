package main

import (
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/http2"
	_ "lets-chat/data"
	"net/http"
)

func main() {

	mux := httprouter.New()
	mux.ServeFiles("/public/*filepath", http.Dir("public"))

	mux.GET("/", Index)

	mux.GET("/signup", Signup)
	mux.POST("/signup_account", SignupAccount)
	mux.GET("/login", Login)
	mux.POST("/authenticate", Authenticate)
	mux.GET("/logout", Logout)

	mux.GET("/oauth/google", GoogleSignUp)
	mux.GET("/oauth/google/callback", GoogleSignUpCallback)
	mux.GET("/complete_signup", CompleteSignup)

	mux.GET("/thread/new", NewThread)
	mux.GET("/threads/:id", FindThread)
	mux.POST("/thread/create", CreateThread)

	mux.POST("/thread/post", CreatePost)


	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	http2.ConfigureServer(server, &http2.Server{})
	server.ListenAndServe()
}
