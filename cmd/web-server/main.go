package main

import (
	"log"
	"net/http"

	"url-checker/cmd/web-server/routes"
)

func main() {
	port := "8000"
	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, http.Handler(routes.RegisterHandlers()))
	if err != nil {
		log.Println(err)
	}
}
