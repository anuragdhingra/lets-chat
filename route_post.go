package main

import (
	"github.com/julienschmidt/httprouter"
	"lets-chat/data"
	"log"
	"net/http"
	"strconv"
)
type PostInfoPublic struct {
	Post data.Post
	CreatedBy data.User
}
func CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sess, err := session(w, r)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
	} else {
		user, _ := sess.User()

		err = r.ParseForm()
		throwError(err)
		postBody := r.PostFormValue("body")
		threadIdString := r.PostFormValue("id")
		threadId, err := strconv.Atoi(threadIdString)
		throwError(err)

		postRequest := data.PostRequest{postBody, user.Id, threadId}
		_, err = postRequest.CreatePost()
		throwError(err)
		http.Redirect(w, r, "/threads/" + threadIdString, http.StatusFound)
	}
}

func CreatePostList(posts []data.Post) (postListPublic []PostInfoPublic) {
	for _, post := range posts {
		postUserId := post.UserId
		user, err := data.UserById(postUserId)
		throwError(err)
		postInfoPublic := PostInfoPublic{post,user}
		postListPublic = append(postListPublic, postInfoPublic)
	}
	return
}