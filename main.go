package main

import (
	_ "github.com/anuragdhingra/lets-chat/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func main() {

	mux := httprouter.New()
	mux.ServeFiles("/public/*filepath", http.Dir("public"))

	mux.GET("/", Index)

	// Auth handlers
	mux.GET("/signup", Signup)
	mux.POST("/signup_account", SignupAccount)
	mux.GET("/login", Login)
	mux.POST("/authenticate", Authenticate)
	mux.GET("/logout", Logout)

	// oauth handlers
	mux.GET("/oauth/google", GoogleSignUp)
	mux.GET("/oauth/google/callback", GoogleSignUpCallback)

	// Thread related handlers
	mux.GET("/thread/new", NewThread)
	mux.GET("/threads/:id", FindThread)
	mux.POST("/thread/create", CreateThread)

	// Post related handlers
	mux.POST("/thread/post", CreatePost)

	server := &http.Server{
		Addr:    "0.0.0.0:" + os.Getenv("port"),
		Handler: mux,
	}

	server.ListenAndServe()
}
