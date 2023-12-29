package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// FIXME: you have to set the StatusCode on the response
// FIXME: you can write a sample text on the response body. No need to use JSON here.
// You can refer to these links to better understand how to propery set a web server up in Go.
// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.golinuxcloud.com/golang-web-server/
// there are more things than you need but try to stick to those implementations since they're the de-facto standards when you need to build a web server by relying only on the Go Std Lib.

func SuccesfulHealthCheck(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, http.StatusOK, r.URL.Path)
}

func BrokenHealthCheck(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, http.StatusInternalServerError, r.URL.Path)
}

func SlowHealthCheck(w http.ResponseWriter, r *http.Request) {
	secondsToWait, err := time.ParseDuration(r.FormValue("wait") + "s")
	if err != nil {
		log.Println("Error on converting seconds to wait", err)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	}
	log.Println("Wait", secondsToWait)
	time.Sleep(secondsToWait)
	sendResponse(w, http.StatusOK, r.URL.Path)
}

func sendResponse(w http.ResponseWriter, response int, api string) {
	log.Println(api, "response with:", response)
	json.NewEncoder(w).Encode(response)
}
