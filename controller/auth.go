package controller

import (
	"net/http"
	"shortyurl/config"
)

func HandleAuth(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, config.ENV.AUTH_URL, http.StatusPermanentRedirect)
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) {

}