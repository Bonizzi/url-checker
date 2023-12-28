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
	// TODO: inspect error scenarios. When an error is triggered?
	response, err := httpClient.Get(url)
	if err != nil {
		urlCheck.Status = response.Status
		return urlCheck
	}
	defer response.Body.Close()
	urlCheck.Status = response.Status
	return urlCheck
}
