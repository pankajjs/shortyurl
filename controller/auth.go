package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"shortyurl/config"
	"shortyurl/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


var (
	oauthConfig = &oauth2.Config{
		ClientID: config.Env.CLIENT_ID,
		ClientSecret: config.Env.CLIENT_SECRET,
		Endpoint: google.Endpoint,
		Scopes: []string{"email", "profile"},
		RedirectURL: config.Env.CALLBACK_URL,
	}
)

type AuthUser struct {
	Name string
	Email string
	Id string
	Verified_email bool
}

var stateToken string

func HandleGoogleAuth(w http.ResponseWriter, r *http.Request) {
	stateToken = utils.GetRandomUuid()
	url:= oauthConfig.AuthCodeURL(stateToken, oauth2.AccessTypeOnline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleAuthCallback(w http.ResponseWriter, r *http.Request) {
	code:= r.URL.Query().Get("code")
	state:= r.URL.Query().Get("state")
	log.Println(code)
	log.Println(state, stateToken, state == stateToken)

	token, err:= oauthConfig.Exchange(context.Background(), code)

	if err != nil {
		return
	}
	
	client:= oauthConfig.Client(context.Background(), token)
	res, err:= client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	
	if err != nil {
		return
	}

	defer res.Body.Close()
	authUser:= new(AuthUser)
	
	if err:= json.NewDecoder(res.Body).Decode(authUser); err != nil {
		return
	}
	
	utils.WriteJson(w, http.StatusOK, authUser)
}