package logic

import (
	"net/http"
	"time"

	"url-checker/models"
)

func CheckStatusAsync(url string, ch chan models.DomainStatus) {
	// FIXME: merge the two following lines into one
	var urlCheck models.DomainStatus
	urlCheck = CheckDomainStatus(url)
	ch <- urlCheck
}

func CheckDomainStatus(url string) models.DomainStatus {
	var urlCheck models.DomainStatus
	urlCheck.Url = url
	httpClient := http.Client{Timeout: 5 * time.Second}
	response, err := httpClient.Get(url)
	if err != nil {
		// FIXME: maybe, it's better to report the status code returned by the hostname invoked. Should be 408
		urlCheck.Status = "Timeout"
		return urlCheck
	}
	defer response.Body.Close()
	urlCheck.Status = response.Status
	return urlCheck
}
