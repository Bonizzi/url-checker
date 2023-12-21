package logic

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"url-checker/models"
)

func AppendToFile(urlCheck models.DomainStatus, split *bool) {
	var file *os.File
	if *split {
		date := time.Now()
		var err error
		domainNoProtocol := strings.SplitAfter(urlCheck.Url, "//")
		domainNoProtocol[1] = strings.ReplaceAll(domainNoProtocol[1], ".", "_")
		file, err = os.OpenFile(date.Format("2006_01_02") + "_" + domainNoProtocol[1] + ".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		ErrorCheck(err)
		defer file.Close()
	} else {
		var err error
		file, err = os.OpenFile("url_list_status.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		ErrorCheck(err)
		defer file.Close()
	}
	w := bufio.NewWriter(file)
	_, err := w.WriteString(urlCheck.Url + " " + urlCheck.Status + "\n")
	ErrorCheck(err)
	w.Flush()
	fmt.Println(urlCheck.Url, "- Checked and saved")
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}