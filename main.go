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

// FIXME: go run . -async -split hangs
// FIXME: add folder /tmp or /out add it to the .gitignore file, otherwise git keeps complaining about it
// TODO: each log entry has to begin with the timestamp formatted in the following way: yyyy-MM-dd hh:mm:urlStatusCh
// TODO: if you split the result through multiple files, be sure to create a folder for the day with the format yyyy_MM_dd
// TODO: the name of the file should be prettied (e.g. no "com", "www", and so on. Just "amazon", "baidu")
// FIXME: manage empty lines within the url.txt file
func main() {
	// FIXME: this should be declared only if we ran the program in the -async mode
	urlStatusCh := make(chan models.DomainStatus, 60)

	// FIXME: go best practices state that you need to use the camelCase notation, not the snake_case.
	async_execution := flag.Bool("async", false, "run in async mode")
	split_file := flag.Bool("split", false, "create a file for each domain")
	flag.Parse()

	// FIXME: there is a bunch of duplicated logic. Try to extrapolate it.
	if *async_execution {
		fmt.Println("Start async mode with split file", *split_file)
		file, err := os.Open("url.txt")
		logic.ErrorCheck(err)
		defer file.Close()
		urlList := bufio.NewScanner(file)
		// FIXME: the time should be recorded at the start.
		t := time.Now()
		for urlList.Scan() {
			go logic.CheckStatusAsync(urlList.Text(), urlStatusCh)
		}
		for i := 0; i < 60; i++ {
			// BUG: use a simple channel instead of the "select" if you have a single branch
			select {
			case msg := <-urlStatusCh:
				logic.AppendToFile(msg, split_file)
			}
		}
		// FIXME: this call should be deferred at the start of the main() func. In this way you're not taking into consideration any function deferred so far.
		fmt.Println("Finished in: ", time.Since(t))
	} else {
		fmt.Println("Start sync mode with split file ", *split_file)
		file, err := os.Open("url.txt")
		logic.ErrorCheck(err)
		defer file.Close()
		urlList := bufio.NewScanner(file)
		// FIXME: the time should be recorded at the start.
		t := time.Now()
		for urlList.Scan() {
			result := logic.CheckStatus(urlList.Text())
			logic.AppendToFile(result, split_file)
		}
		// FIXME: this call should be deferred at the start of the main() func. In this way you're not taking into consideration any function deferred so far.
		fmt.Println("Finished in: ", time.Since(t))
	}
}
