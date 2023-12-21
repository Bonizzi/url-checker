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
	urlStatusCh := make(chan models.DomainStatus, 60)

	async_execution := flag.Bool("async", false, "run in async mode")
	split_file := flag.Bool("split", false, "create a file for each domain")
	flag.Parse()

	if *async_execution {
		fmt.Println("Start async mode with split file", *split_file)
		file, err := os.Open("url.txt")
		logic.ErrorCheck(err)
		defer file.Close()
		urlList := bufio.NewScanner(file)
		t := time.Now()
		for urlList.Scan() {
			go logic.CheckStatusAsync(urlList.Text(), urlStatusCh)
		}
		for i := 0; i < 60; i++ {
			select {
			case msg := <- urlStatusCh:
				logic.AppendToFile(msg, split_file)
			}
		}
		fmt.Println("Finished in: ", time.Since(t))
	} else {
		fmt.Println("Start sync mode with split file ", *split_file)
		file, err := os.Open("url.txt")
		logic.ErrorCheck(err)
		defer file.Close()
		urlList := bufio.NewScanner(file)
		t := time.Now()
		for urlList.Scan() {
			result := logic.CheckStatus(urlList.Text())
			logic.AppendToFile(result, split_file)
		}
		fmt.Println("Finished in: ", time.Since(t))
	}
}