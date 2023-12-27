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
		os.Mkdir(*folderPath, 0o700)
	}
	if *split {
		defaultDailyFolder := time.Now().Format("2006_01_02")
		if !doesExist(*folderPath+"/"+defaultDailyFolder, "") {
			os.Mkdir(*folderPath+"/"+defaultDailyFolder, 0o700)
		}
		// FIXME: try using filepath pkg of the Std Lib
		filePath := *folderPath + "/" + defaultDailyFolder
		var err error
		// FIXME: extrapolate this logic into a function that accepts a string input param and return another string value, plus an error if something is wrong.
		domainNoProtocol := strings.SplitAfter(urlCheck.Url, "www.")
		justDomainName := strings.Split(domainNoProtocol[1], ".")
		fileName := defaultDailyFolder + "_" + justDomainName[0] + ".txt"
		if doesExist(filePath, fileName) {
			fileName = strings.ReplaceAll(fileName, justDomainName[0]+".txt", strings.ReplaceAll(domainNoProtocol[1], ".", "_")+".txt")
		}
		file, err = os.OpenFile(filePath+"/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		ErrorCheck(err)
		defer file.Close()
	} else {
		var err error
		// FIXME: add to the .gitignore every file called "url_list_status.txt" located anywhere
		file, err = os.OpenFile(*folderPath+"/"+"url_list_status.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		ErrorCheck(err)
		defer file.Close()
	}
	_, err := file.WriteString(urlCheck.Url + " " + urlCheck.Status + "\n")
	ErrorCheck(err)
	// [Q]: do you need this Sync method? If so, you have to check the error.
	file.Sync()
	PrintLog(urlCheck.Url + " - Checked and saved")
}

// BUG: avoid this func. If you encounter an error you should immediately STOP the execution. Otherwise, you can get a nil dereference error.
func ErrorCheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PrintLog(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", msg)
}

func GetFileNumLines(file *os.File) int {
	// FIXME: not meaningful name
	urlList := bufio.NewScanner(file)
	count := 0
	for urlList.Scan() {
		if urlList.Text() != "" {
			count++
		}
	}
	// BUG: always checks for urlList.Err() and return it
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		// BUG: return the error to be handled above in the chain
		log.Fatal(err)
	}
	return count
}

func doesExist(folderPath string, fileName string) bool {
	// FIXME: filepath pkg
	path := folderPath + "/" + fileName
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

func GetLocalPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	return basepath
}
