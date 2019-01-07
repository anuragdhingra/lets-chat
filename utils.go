package main

import (
	"chit-chat/data"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func generateHTML(writer http.ResponseWriter, threads []data.Thread, filenames ...string) {
	 var files []string
	 for _, file := range filenames {
	 	files = append(files, fmt.Sprintf("templates/%s.html", file))
	 }

	 templates := template.Must(template.ParseFiles(files...))
	 templates.ExecuteTemplate(writer,"layout", threads)
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