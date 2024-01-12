package controller

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// 'https://www.golinuxcloud.com/golang-web-server/'

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	switch method := r.URL; fmt.Sprint(method) {
	case "/health":
		if checkGetMethod(w, r, fmt.Sprint(method)) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprint(http.StatusOK)))
		}
	case "/broken":
		if checkGetMethod(w, r, fmt.Sprint(method)) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(http.StatusInternalServerError)))
		}
	default:
		if checkGetMethod(w, r, fmt.Sprint(method)) {
			SlowHealthCheck(w, r)
		}
	}
}

func SlowHealthCheck(w http.ResponseWriter, r *http.Request) {
	if regexp.MustCompile(`^[0-9]*$`).MatchString(r.URL.Query().Get("wait")) && len(r.URL.Query().Get("wait")) > 0 {
		secondsToWait, err := time.ParseDuration(r.URL.Query().Get("wait") + "s")
		if err != nil {
			log.Println("ERROR [SlowHealthCheck] Error on converting seconds to wait", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprint(http.StatusInternalServerError)))
		}
		log.Println("INFO [SlowHealthCheck] Wait", secondsToWait)
		time.Sleep(secondsToWait)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprint(http.StatusOK)))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint(http.StatusBadRequest)))
		log.Println("ERROR [SlowHealthCheck] Invalid parameter received")
	}
}

func checkGetMethod(w http.ResponseWriter, r *http.Request, endpoint string) bool {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprint(http.StatusMethodNotAllowed)))
		return false
	}
	return true
}
