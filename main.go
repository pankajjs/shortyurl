package main

import (
	"log"
	"shortyurl/cmd/api"
	"shortyurl/config"
)

func main(){
	port:= config.ENV.PORT
	api:= api.NewApiServer(port)

	log.Printf("Attempting to start server at port %v", port)
	
	if err:= api.Run(); err != nil {
		log.Fatalf("Server failed to start at port %v, reason: %v", port, err)
	}
}