package main

import (
	"chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	generateHTML(w, nil, "layout", "public.navbar", "signup")
}

func SignupAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)

	user := data.User{
		Username:r.PostFormValue("username"),
		Email:r.PostFormValue("email"),
		Password: encryptPassword(r.PostFormValue("password")),
	}

	err = user.Create()
	throwError(err)
	http.Redirect(w, r, "/", 302)
}