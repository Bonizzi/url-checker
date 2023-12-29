package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
	"url-checker/cmd/url-check/internal/logic"
	"url-checker/cmd/url-check/models"
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
	if err != nil {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", "ERROR", err)
		return
	}
	defer file.Close()
	urlList := bufio.NewScanner(file)
	if *asyncExecution {
		lines, err := logic.GetFileNumLines(file)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", "ERROR:", err)
			return
		}
		urlStatusCh := make(chan models.DomainStatus, lines)
		for urlList.Scan() {
			if urlList.Text() != "" {
				go logic.CheckStatusAsync(urlList.Text(), urlStatusCh)
			}
		}
		err2 := urlList.Err()
		if err2 != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", "Error listing in async mode", err2)
			return
		}
		for i := 0; i < lines; i++ {
			msg := <-urlStatusCh
			err := logic.AppendToFile(msg, splitFile, outputFileName)
			if err != nil {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", err)
			}
		}
	} else {
		for urlList.Scan() {
			if urlList.Text() != "" {
				result := logic.CheckDomainStatus(urlList.Text())
				err := logic.AppendToFile(result, splitFile, outputFileName)
				if err != nil {
					fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", err)
				}
			}
		}
		err := urlList.Err()
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05")+":", "Error listing in sync mode", err)
			return
		}
	}
}
