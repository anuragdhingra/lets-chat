package main

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"lets-chat/data"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL: "http://localhost:8080/login/google/callback",
	ClientID: os.Getenv("clientId"),
	ClientSecret: os.Getenv("clientSecret"),
	Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

type GoogleUserInfo struct {
	Id int `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Link url.URL `json:"link"`
	Picture url.URL `json:"picture"`
}

var cookieee http.Cookie

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GoogleSignUp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u4, _ := uuid.NewV4()
	oauthState := u4.String()
	cookie := http.Cookie{
		Name:"oauthState",
		Value:oauthState,
		HttpOnly: true,
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
	http.SetCookie(w, &cookie)

	url := googleOauthConfig.AuthCodeURL(oauthState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleSignUpCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	oauthStateCookie, _ := r.Cookie("oauthState")
	oauthState := oauthStateCookie.Value
	state := r.URL.Query().Get("state")
	code := r.URL.Query().Get("code")

	if state != oauthState {
		log.Print("Invalid oauth state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	throwError(err)

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	throwError(err)

	defer response.Body.Close()

	userInfoBytes, err := ioutil.ReadAll(response.Body)
	throwError(err)

	userInfo := &GoogleUserInfo{}
	_ = json.Unmarshal(userInfoBytes, &userInfo)

	username := strings.Split(userInfo.Email, "@")[0]

	// Create user if it doesn't exists
	_, err = data.UserByEmailOrUsername(userInfo.Email)
	if err != nil {
		user := data.User{
			Username:    username,
			Email:       userInfo.Email,
			HasPassword: false,
		}

		_ = user.Create()
	}

	// Log in the user
	loggedInUser, err := data.UserByEmailOrUsername(userInfo.Email)
	throwError(err)
	session, err := loggedInUser.CreateSession()
	cookieee = http.Cookie{
		Name:"_cookie",
		Value: session.Uuid,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookieee)
	http.Redirect(w, r, "/complete_signup", http.StatusTemporaryRedirect)

}

func CompleteSignup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.SetCookie(w, &cookieee)
	_, err := session(w, r)
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
	}
}