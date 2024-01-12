package main

import (
	"errors"
	"log"
	"net/http"

	"url-checker/cmd/web-server/routes"

	"github.com/rs/cors"
)

// FIXME: no need to use CORS here. This should be used to build SPA (client + server)

func main() {
	port := "8000"
	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, cors.AllowAll().Handler(routes.StartRoutes()))
	if err != nil {
		// FIXME: "Unwrap" must not be used here. You're already at the top-level of your app.
		// And there is no wrapped error to handle since the it's generated here.
		log.Println("[ERROR] Port already in use:", errors.Unwrap(errors.Unwrap(err)))
	}
}
