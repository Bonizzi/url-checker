package logic

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"url-checker/cmd/url-check/models"
)

func CheckStatusAsync(url string, ch chan models.DomainStatus) {
	urlCheck := CheckDomainStatus(url)
	ch <- urlCheck
}

func CheckDomainStatus(url string) models.DomainStatus {
	var urlCheck models.DomainStatus
	urlCheck.Url = url
	var httpClient http.Client
	if !strings.Contains(url, "slowdown") { // condition to handle the /slowdown request which already contains a timeout
		httpClient = http.Client{Timeout: 5 * time.Second}
	}
	response, err := httpClient.Get(url)
	if err != nil {
		fmt.Println("Http GET Error", err)
		urlCheck.Status = http.StatusRequestTimeout
		return urlCheck
	}
	defer response.Body.Close()
	urlCheck.Status = response.StatusCode
	return urlCheck
}
