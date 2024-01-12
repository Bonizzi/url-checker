package routes

import (
	"url-checker/cmd/web-server/controller"

	"github.com/gorilla/mux"
)

// FIXME: you have two options:
// 1. delete this package and instrument the routes within the "main" func
// 2. rename this func to "RegisterHandlers" since the name is not meaningful.
func StartRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", controller.SuccesfulHealthCheck)
	r.HandleFunc("/slowdown", controller.SlowHealthCheck)
	r.HandleFunc("/broken", controller.BrokenHealthCheck)

	return r
}
