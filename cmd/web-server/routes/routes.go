package routes

import (
	"url-checker/cmd/web-server/controller"

	"github.com/gorilla/mux"
)

func RegisterHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", controller.HealthCheck)
	r.HandleFunc("/slowdown2", controller.SlowHealthCheck2)
	r.HandleFunc("/slowdown", controller.HealthCheck)
	r.HandleFunc("/broken", controller.HealthCheck)

	return r
}
