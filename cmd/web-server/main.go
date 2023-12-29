package main

import (
	"errors"
	"log"
	"net/http"

	"url-checker/cmd/web-server/routes"

	"github.com/rs/cors"
)

func main() {
	port := "8000"
	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, cors.AllowAll().Handler(routes.StartRoutes()))
	if err != nil {
		log.Println("[ERROR] Port already in use:", errors.Unwrap(errors.Unwrap(err)))
	}
}
