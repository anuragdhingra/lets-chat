package main

import (
	"chit-chat/data"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		log.Print(err)
		return
	} else {
		log.Print("index fired")
		generateHTML(w, threads, "layout", "public.navbar", "index")
	}
}