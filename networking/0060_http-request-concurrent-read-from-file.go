package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sendRequest(url string) {
	defer wg.Done()
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%d] %s\n", res.StatusCode, url)
}

func main() {
	start := time.Now()

	readFile, err := os.Open("0060_urls.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	readFile.Close()

	for _, url := range fileTextLines {
		go sendRequest("https://" + url)
		// sendRequest("https://" + url)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Printf("\n%s\n", time.Since(start))
}

// 10.77
// 1.67
