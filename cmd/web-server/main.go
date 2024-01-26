package main

import (
	"log"
	"net/http"

	"url-checker/cmd/web-server/routes"
)

// BUG: issue with the elapsed time.
// To reproduce it, I changed this call in the "url.txt" file:
// old: http://localhost:8000/slowdown?wait=5
// new: http://localhost:8000/slowdown?wait=7
// when I run it async, it gives me ~5s as elapsed time.
// when I run this: curl -o /dev/null -s -w 'Total: %{time_total}s\n' http://localhost:8000/slowdown?wait=7
// it gives me ~7s which is OK.
func main() {
	port := "8000"
	log.Println("Starting server on port " + port)
	err := http.ListenAndServe(":"+port, http.Handler(routes.RegisterHandlers()))
	if err != nil {
		log.Println(err)
	}
}
