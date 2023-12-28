package routes

import (
	"url-checker/web-server/controller"

	"github.com/gorilla/mux"
)

func StartRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/health", controller.SuccesfulHealthCheck)
	r.HandleFunc("/slowdown", controller.SlowHealthCheck)
	r.HandleFunc("/broken", controller.BrokenHealthCheck)

	return r
}
