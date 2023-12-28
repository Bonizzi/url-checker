package logic

import (
	"net/http"
	"time"

	"url-checker/models"
)

func CheckStatusAsync(url string, ch chan models.DomainStatus) {
	urlCheck := CheckDomainStatus(url)
	ch <- urlCheck
}

func CheckDomainStatus(url string) models.DomainStatus {
	var urlCheck models.DomainStatus
	urlCheck.Url = url
	httpClient := http.Client{Timeout: 5 * time.Second}
	response, err := httpClient.Get(url)
	// TODO: do something with the `err` variable
	if err != nil {
		urlCheck.Status = http.StatusRequestTimeout
		return urlCheck
	}
	defer response.Body.Close()
	urlCheck.Status = response.StatusCode
	return urlCheck
}
