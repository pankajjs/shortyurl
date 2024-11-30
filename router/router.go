package router

import (
	"net/http"
	"shortyurl/controller"
	"shortyurl/utils"
)

func NewMainRouter() *http.ServeMux {
	router:= http.NewServeMux()
	router.Handle("/", NewPageRouter())
	router.Handle("/api/", http.StripPrefix("/api", NewApiRouter()))
	router.HandleFunc("/healthcheck", controller.Healthcheck)
	return router
}

func NewPageRouter() *http.ServeMux {
	router:= http.NewServeMux()
	return router
}

func NewApiRouter() *http.ServeMux {
	router:= http.NewServeMux()
	router.Handle("/v1/", http.StripPrefix("/v1", NewV1Router()))
	return router
}

func NewV1Router() *http.ServeMux {
	router:= http.NewServeMux()
	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJson(w, http.StatusOK, "Users")
	})
	return router
}