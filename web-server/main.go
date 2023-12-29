package main

import (
	"log"
	"net/http"

	"url-checker/web-server/routes"

	"github.com/rs/cors"
)

// FIXME: use the "cmd" structure. Put the executables below child folders of "cmd" one.
// refer to this: https://go.dev/doc/modules/layout

// FIXME: some broken calls:
/*
curl -X POST http://localhost:8000/health (the endpoint must be invokable with only the GET method)
curl http://localhost:8000/slowdown?waits=5 (wrong query param name)
*/

// TODO: add these urls to the url.txt
func main() {
	port := "8000"
	log.Println("Starting server on port " + port)
	// FIXME: always checks for the error. I can try to run two servers on the same port. With the second try you should tell me what was wrong (e.g. port already allocated).
	http.ListenAndServe(":"+port, cors.AllowAll().Handler(routes.StartRoutes()))
}
