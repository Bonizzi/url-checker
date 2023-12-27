package logic

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"url-checker/models"
)

func AppendToFile(urlCheck models.DomainStatus, split *bool, folderPath *string) {
	var file *os.File
	if !doesExist(*folderPath, "") {
		os.Mkdir(*folderPath, 0700)
	}
	if *split {
		defaultDailyFolder := time.Now().Format("2006_01_02")
		if !doesExist(*folderPath+"/"+defaultDailyFolder, "") {
			os.Mkdir(*folderPath+"/"+defaultDailyFolder, 0700)
		}
		filePath := *folderPath+"/"+defaultDailyFolder
		var err error
		domainNoProtocol := strings.SplitAfter(urlCheck.Url, "www.")
		justDomainName := strings.Split(domainNoProtocol[1], ".")
		fileName := defaultDailyFolder+"_"+justDomainName[0]+".txt"
		if doesExist(filePath, fileName) {
			fileName = strings.ReplaceAll(fileName, justDomainName[0] + ".txt", strings.ReplaceAll(domainNoProtocol[1], ".", "_") + ".txt")
		}
		file, err = os.OpenFile(filePath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		ErrorCheck(err)
		defer file.Close()
	} else {
		var err error
		file, err = os.OpenFile(*folderPath+"/"+"url_list_status.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		ErrorCheck(err)
		defer file.Close()
	}
	_, err := file.WriteString(urlCheck.Url + " " + urlCheck.Status + "\n")
	ErrorCheck(err)
	file.Sync()
	PrintLog(urlCheck.Url + " - Checked and saved")
}

func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PrintLog(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + ":",msg)
}

func GetFileNumLines(file *os.File) int {
	urlList := bufio.NewScanner(file)
	count := 0
	for urlList.Scan() {
		if urlList.Text() != "" {
			count++
		}
	}
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func doesExist(folderPath string, fileName string) bool {
	path := folderPath+"/"+fileName
	if fileName == "" {
		path = folderPath
	}
	_ , error := os.Stat(path)
	if os.IsNotExist(error) {
	  return false
	} else {
	  return true
	}
  }

func GetLocalPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	return basepath
}
