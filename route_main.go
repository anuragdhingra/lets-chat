package main

import (
	"chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	threads, err := data.Threads()
	if err != nil {
		log.Print(err)
		return
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, threads, "layout","public.navbar", "index")
		} else {
			generateHTML(w, threads, "layout", "private.navbar","index")
		}
	}
	}

func FindThread(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	threadId := p.ByName("id")
	thread, err := data.ThreadByID(threadId)
	throwError(err)
	if err != nil {
		log.Print(err)
		return
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, thread, "layout","public.navbar", "thread")
		} else {
			generateHTML(w, thread, "layout", "private.navbar","thread")
		}
	}
}