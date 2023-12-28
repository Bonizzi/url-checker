package main

import (
	"log"
	"net/http"
	"url-checker/web-server/routes"

	"github.com/rs/cors"
)

func main() {
	port := "8000"
	log.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, cors.AllowAll().Handler(routes.StartRoutes()))
}
