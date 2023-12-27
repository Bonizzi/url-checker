package logic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"url-checker/models"
)

func AppendToFile(urlCheck models.DomainStatus, split *bool, outputFileName *string) {
	var file *os.File
	folder := "tmp"
	if !doesExist(folder, "") {
		os.Mkdir(folder, 0o700)
	}
	if *split {
		defaultDailyFolder := time.Now().Format("2006_01_02")
		filePath := filepath.Join(folder, defaultDailyFolder)
		if !doesExist(filePath, "") {
			os.Mkdir(filePath, 0o700)
		}
		var err error
		file, err = os.OpenFile(filepath.Join(filePath, createFileName(urlCheck.Url, defaultDailyFolder, filePath)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	} else {
		var err error
		file, err = os.OpenFile(filepath.Join(folder, checkFileExtention(outputFileName)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
	}
	_, err := file.WriteString(time.Now().Format("2006-01-02 15:04") + " " + urlCheck.Url + " " + urlCheck.Status + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	PrintLog(urlCheck.Url + " - Checked and saved")
}

func PrintLog(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", msg)
}

func GetFileNumLines(file *os.File) int {
	listOfAllUrls := bufio.NewScanner(file)
	count := 0
	for listOfAllUrls.Scan() {
		if listOfAllUrls.Text() != "" {
			count++
		}
	}
	err := listOfAllUrls.Err()
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := file.Seek(0, io.SeekStart)
	if err2 != nil {
		fmt.Println(err2)
	}
	return count
}

func doesExist(folderPath string, fileName string) bool {
	path := filepath.Join(folderPath, fileName)
	if fileName == "" {
		path = folderPath
	}
	_, error := os.Stat(path)
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

func checkFileExtention(fileName *string) string {
	if filepath.Ext(*fileName) == "" {
			*fileName += ".txt"
	}
	return *fileName
}

func createFileName(url string, defaultDailyFolder string, filePath string) string {
	domainNoProtocol := strings.SplitAfter(url, "www.")
	justDomainName := strings.Split(domainNoProtocol[1], ".")
	fileName := defaultDailyFolder + "_" + justDomainName[0] + ".txt"

	if doesExist(filePath, fileName) {
		fileName = strings.ReplaceAll(fileName, justDomainName[0]+".txt", strings.ReplaceAll(domainNoProtocol[1], ".", "_")+".txt")
	}
	return fileName
}