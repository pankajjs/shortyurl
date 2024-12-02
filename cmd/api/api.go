package api

import (
	"net/http"
	"shortyurl/router"
)


type ApiServer struct {
	Addr string
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		Addr: addr,
	}
}

func (a *ApiServer) Run() error {	
	return http.ListenAndServe(a.Addr, router.ApiRouter())
}