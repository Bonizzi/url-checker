package logic

import (
	"fmt"
	"net/http"
	"time"
	"url-checker/models"
)

func CheckStatus(url string) models.DomainStatus {
	var urlCheck models.DomainStatus
	httpClient := http.Client{Timeout: 5 * time.Second}
	response, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(url, "Timeout")
		urlCheck.Url = url
		urlCheck.Status = "Timeout"
		return urlCheck
	}
	defer response.Body.Close()
	urlCheck.Url = url
	urlCheck.Status = response.Status
	return urlCheck
}

func CheckStatusAsync(url string, ch chan models.DomainStatus) {
	var urlCheck models.DomainStatus
	httpClient := http.Client{Timeout: 5 * time.Second}
	response, err := httpClient.Get(url)
	if err != nil {
		fmt.Println(url, "Timeout")
		urlCheck.Url = url
		urlCheck.Status = "Timeout"
		ch <- urlCheck
	}
	defer response.Body.Close()
	urlCheck.Url = url
	urlCheck.Status = response.Status
	ch <- urlCheck
}