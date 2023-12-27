package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"

	"url-checker/internal/logic"
	"url-checker/models"
)

// FIXME: by default, every output file must be written below a folder "tmp/". You can choose the name of the file but the location must be within this folder.
// FIXME: each log entry (in the files) has to begin with the timestamp formatted in the following way: yyyy-MM-dd hh:mm
func main() {
	asyncExecution := flag.Bool("async", false, "Run in async mode")
	splitFile := flag.Bool("split", false, "Create a file for each domain")
	folderPath := flag.String("path", logic.GetLocalPath(), "To specify a folder path to save the output file(s)")
	flag.Parse()

	t := time.Now()
	defer func() {
		fmt.Println("Finished in:", time.Since(t))
	}()
	logic.PrintLog("START URL CHECK")
	file, err := os.Open("url.txt")
	logic.ErrorCheck(err)
	defer file.Close()
	urlList := bufio.NewScanner(file)
	if *asyncExecution {
		lines := logic.GetFileNumLines(file)
		urlStatusCh := make(chan models.DomainStatus, lines)
		for urlList.Scan() {
			if urlList.Text() != "" {
				go logic.CheckStatusAsync(urlList.Text(), urlStatusCh)
			}
		}
		// BUG: always checks for urlList.Err() and return it
		for i := 0; i < lines; i++ {
			msg := <-urlStatusCh
			logic.AppendToFile(msg, splitFile, folderPath)
		}
	} else {
		for urlList.Scan() {
			if urlList.Text() != "" {
				result := logic.CheckDomainStatus(urlList.Text())
				logic.AppendToFile(result, splitFile, folderPath)
			}
		}
		// BUG: always checks for urlList.Err() and return it
	}
}
