package main

import (
	"chit-chat/data"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	 var files []string
	 for _, file := range filenames {
	 	files = append(files, fmt.Sprintf("templates/%s.html", file))
	 }

	 templates := template.Must(template.ParseFiles(files...))
	 templates.ExecuteTemplate(writer,"layout", data)
}

func throwError(err error) {
	if err != nil {
		log.Print(err)
		return
	}
}

func encryptPassword(password string) (encryptedPass string) {
	bytePass := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.MinCost)
	throwError(err)
	encryptedPass = string(encryptedPassword)

	return string(encryptedPass)
}

func session(w http.ResponseWriter, r *http.Request) (session data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Print(err)
		return
	} else {
		session = data.Session{Uuid:cookie.Value}
		ok,_ := session.Check()
		if !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func checkInvalidRequests(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		log.Print("Session token not found")
		http.Redirect(w, r, "/", 302)
		return
	}
}