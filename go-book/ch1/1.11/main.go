package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var fClient = &http.Client{
	Timeout: 10 * time.Second,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:\n\t./1.11 <infile>\n")
		os.Exit(1)
	}
	start := time.Now()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ch := make(chan string)
	pool := 0
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		url := fmt.Sprintf("https://%s", record[1])
		go fetch(url, ch)
		pool++
	}

	for pool > 0 {
		fmt.Println(<-ch)
		pool--
	}
	fmt.Printf("Finished in %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := fClient.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("failed to get url %q: %v", url, err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("failed to read body: %v", err)
		return
	}
	ch <- fmt.Sprintf("%s\tfetched %d bytes in %.2fs", url, nbytes, time.Since(start).Seconds())
}
