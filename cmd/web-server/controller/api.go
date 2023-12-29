package controller

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
// https://www.golinuxcloud.com/golang-web-server/

func SuccesfulHealthCheck(w http.ResponseWriter, r *http.Request) {
	if !checkGetMethod(w, r) {
		log.Println("ERROR [SuccesfulHealthCheck] Invalid method received:", r.Method)
	} else {
		sendResponse(w, http.StatusOK, r)
	}
}

func BrokenHealthCheck(w http.ResponseWriter, r *http.Request) {
	if !checkGetMethod(w, r) {
		log.Println("ERROR [BrokenHealthCheck] Invalid method received:", r.Method)
	} else {
		sendResponse(w, http.StatusInternalServerError, r)
	}
}

func SlowHealthCheck(w http.ResponseWriter, r *http.Request) {
	if checkGetMethod(w, r) {
		if regexp.MustCompile(`^[0-9]*$`).MatchString(r.URL.Query().Get("wait")) && len(r.URL.Query().Get("wait")) > 0 {
			secondsToWait, err := time.ParseDuration(r.URL.Query().Get("wait") + "s")
			if err != nil {
				log.Println("ERROR [SlowHealthCheck] Error on converting seconds to wait", err)
				io.WriteString(w, strconv.Itoa(http.StatusInternalServerError))
			}
			log.Println("INFO [SlowHealthCheck] Wait", secondsToWait)
			time.Sleep(secondsToWait)
			sendResponse(w, http.StatusOK, r)
		} else {
			io.WriteString(w, "Invalid or missing wait parameter")
			log.Println("ERROR [SlowHealthCheck] Invalid parameter received")
		}
	} else {
		log.Println("ERROR [SlowHealthCheck] Invalid method received:", r.Method)
	}
}

func sendResponse(w http.ResponseWriter, response int, r *http.Request) {
	log.Println(r.URL.Path, "response with:", response)
	_, err := io.WriteString(w, strconv.Itoa(response))
	if err != nil {
		log.Println("ERROR [sendResponse] Cannot send response", err)
	}
}

func checkGetMethod(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == "POST" {
		_, err := io.WriteString(w, "Method "+r.Method+" not accepted, only GET")
		if err != nil {
			log.Println("ERROR [checkGetMethod] Cannot send invalid method response", err)
		}
		return false
	}
	return true
}
