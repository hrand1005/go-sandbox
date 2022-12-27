package main

import (
    "github.com/hrand1005/go-sandbox/go-book/ch5/links"
    "os"
    "log"
)

func main() {
    worklist := make(chan []string)
    go func() { worklist <- os.Args[1:] }()

    seen := make(map[string]bool)
    for n := 1; n > 0; n-- {
        list := <- worklist
        for _, link := range list {
            if !seen[link] {
                n++
                seen[link] = true
                go func(link string) {
                    worklist <- crawl(link)
                }(link)
            }
        }
    }
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
    tokens <- struct{}{}
    list, err := links.Extract(url)
    <- tokens
    if err != nil {
        log.Println(err)
    }

    return list
}
