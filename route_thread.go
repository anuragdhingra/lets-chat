package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func NewThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess, err := session(w, r)
	if err == nil {
		loggedInUser, err := sess.User()
		data := Data{nil, loggedInUser}
		if err !=nil {
			log.Print(err)
			return
		} else {
			generateHTML(w, data, "layout", "private.navbar", "new.thread")
		}
	} else {
		log.Print(err)
		http.Redirect(w, r, "/login", 302)
		return
	}
}