package router

import (
	"net/http"
	"shortyurl/controller"
	"shortyurl/utils"
)

func ApiRouter() *http.ServeMux {
	router:= http.NewServeMux()
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", V1Router()))
	router.HandleFunc("/healthcheck", controller.Healthcheck)
	return router
}

func V1Router() *http.ServeMux {
	router:= http.NewServeMux()
	router.Handle("/auth/", http.StripPrefix("/auth", AuthRouter()))
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJson(w, http.StatusOK, "Users")
	})
	return router
}

func AuthRouter() *http.ServeMux {
	router:= http.NewServeMux()
	router.HandleFunc("/google", controller.HandleAuth)
	router.HandleFunc("/callback/google", controller.HandleAuthCallback)
	return router
}