package main

import (
	"chit-chat/data"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var googleOauthConfig = &oauth2.Config{
RedirectURL: "http://localhost:8080/login/google/callback",
ClientID: "270815731229-eoou8vanhp71jb8d8vjq13fo1l7sggod.apps.googleusercontent.com",
ClientSecret: "1Bu5AcreI3W6AlnqqjBPFxtG",
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

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func Signup(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkInvalidRequests(w, r)
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
	http.Redirect(w, r, "/login", 200)
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	checkInvalidRequests(w,r)
	generateHTML(w, nil, "layout", "public.navbar", "login")
}

func Authenticate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseForm()
	throwError(err)

	emailOrUsername := r.PostFormValue("emailOrUsername")
	pass := r.PostFormValue("password")

	user, err := data.UserByEmailOrUsername(emailOrUsername)
	throwError(err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		log.Print(err)
		http.Redirect(w, r, "/login", 302)
		return
	} else {
		session, err := user.CreateSession()
		throwError(err)

		cookie := http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w,&cookie)
 
		log.Print("User successfully logged in")
		http.Redirect(w, r, "/", 302)
	}
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		sess := data.Session{Uuid:cookie.Value}
		err = sess.DeleteByUUID()
		if err != nil {
			log.Print(err)
			return
		} else {
			cookie := http.Cookie{
				Name:   "_cookie",
				MaxAge: -1,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", 302)
		}
	} else {
		log.Print("Invalid request")
		http.Redirect(w, r, "/", 302)
	}
}

func GoogleSignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func GoogleSignInCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	user := data.User{
		Username: username,
		Email:userInfo.Email,
		Password: encryptPassword("temppassword"),
	}

	err = user.Create()
	throwError(err)
	http.Redirect(w, r, "/login", 302)
}
