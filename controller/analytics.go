package controller

import (
	"net/http"
	"shortyurl/config"
	"shortyurl/utils"
	"time"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"up": time.Since(config.StartTime).String(),
	})
}