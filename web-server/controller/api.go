package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

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
