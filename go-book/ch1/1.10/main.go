package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const outfile = "out.txt"

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	f, err := os.Create(outfile)
	if err != nil {
		log.Fatalf("failed to open %s: %v\n", outfile, err)
	}
	defer f.Close()
	for range os.Args[1:] {
		if _, err := f.WriteString(<-ch + "\n"); err != nil {
			log.Printf("failed to write string: %v\n", err)
		}
	}
	fmt.Printf("Finished in %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
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
