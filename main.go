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

func main() {
	asyncExecution := flag.Bool("async", false, "Run in async mode")
	splitFile := flag.Bool("split", false, "Create a file for each domain")
	outputFileName := flag.String("file", "url_list_status.txt", "To specify a folder path to save the output file(s)")
	flag.Parse()

	t := time.Now()
	defer func() {
		fmt.Println("Finished in:", time.Since(t))
	}()
	logic.PrintLog("START URL CHECK")
	file, err := os.Open("url.txt")
	// TODO: error management
	// https://earthly.dev/blog/golang-errors/
	if err != nil {
		fmt.Println(err)
		return
	}
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
		err := urlList.Err()
		if err != nil {
			fmt.Println(err)
		}
		for i := 0; i < lines; i++ {
			msg := <-urlStatusCh
			logic.AppendToFile(msg, splitFile, outputFileName)
		}
	} else {
		for urlList.Scan() {
			if urlList.Text() != "" {
				result := logic.CheckDomainStatus(urlList.Text())
				logic.AppendToFile(result, splitFile, outputFileName)
			}
		}
		err := urlList.Err()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
