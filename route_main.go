package main

import (
	"chit-chat/data"
	"net/http"
)

func index(w http.ResponseWriter, r http.Request) {
	threads, err := data.Threads()
	if err != nil {
		// throws error
	} else {
		generateHTML(w, threads, "layout", "public.navbar", "index")
	}
}