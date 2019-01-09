package main

import (
	"chit-chat/data"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

type ThreadInfoPublic struct {
	Thread data.Thread
	CreatedBy data.User
}

type ThreadInfoPrivate struct {
	Thread data.Thread
	CreatedBy data.User
	User data.User
}

func NewThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess, err := session(w, r)
	if err == nil {
		loggedInUser, err := sess.User()
		data := ThreadsInfoPrivate{nil, loggedInUser}
		if err != nil {
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

func CreateThread(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)
	sess, err := session(w, r)
	throwError(err)
	user, err := sess.User()
	throwError(err)

	createThreadRequest := data.CreateThreadRequest{
		r.PostFormValue("topic"),
		user.Id,

	}
	threadId, err := createThreadRequest.Create()
	log.Print(threadId)
	throwError(err)
	url := "/threads/" +  strconv.Itoa(threadId)
	log.Print(url)
	http.Redirect(w, r, url, 302)
}

func FindThread(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	threadId := p.ByName("id")
	thread, err := data.ThreadByID(threadId)
	if err != nil {
		log.Print(err)
		return
	} else {
		user, err := data.UserById(thread.UserId)
		throwError(err)
		sess, err := session(w, r)
		if err != nil {
			data := ThreadInfoPublic{thread, user}
			generateHTML(w, data, "layout","public.navbar", "thread")
		} else {
			loggedInUser, err := sess.User()
			throwError(err)
			data := ThreadInfoPrivate{thread, user, loggedInUser}
			generateHTML(w, data, "layout", "private.navbar","thread")
		}
	}
}

func CreateThreadList(threads []data.Thread) (threadListPublic []ThreadInfoPublic) {
	for _, thread := range threads {
		threadUserId := thread.UserId
		user, err := data.UserById(threadUserId)
		throwError(err)
		threadInfoPublic := ThreadInfoPublic{thread,user}
		threadListPublic = append(threadListPublic, threadInfoPublic)
	}
	return
}