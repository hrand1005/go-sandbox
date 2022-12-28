// Exercise 8.6
package main

import (
	"flag"
	"github.com/hrand1005/go-sandbox/go-book/ch5/links"
	"log"
	"strings"
)

var startlist = flag.String("startlist", "", "comma separated list of domains to start crawling")
var depth = flag.Int("depth", 3, "depth limit for link traversal")

type linksWithDepth struct {
	links []string
	depth int
}

func main() {
	flag.Parse()
	if *startlist == "" {
		flag.Usage()
		return
	}

	startlist := &linksWithDepth{
		links: strings.Split(*startlist, ","),
		depth: 0,
	}

	worklist := make(chan *linksWithDepth)
	go func() { worklist <- startlist }()

	seen := make(map[string]bool)
	n := 1
	for n > 0 {
		next := <-worklist
		n--
		if next.depth == *depth {
			continue
		}

		for _, link := range next.links {
			if !seen[link] {
				n++
				seen[link] = true
				go func(link string, newDepth int) {
					worklist <- &linksWithDepth{
						links: crawl(link),
						depth: newDepth,
					}
				}(link, next.depth+1)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens
	if err != nil {
		log.Println(err)
	}

	return list
}
