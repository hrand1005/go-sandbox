package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineData struct {
	line     string
	filename string
}

func main() {
	counts := make(map[lineData]int)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}
	for ld, n := range counts {
		if n > 1 {
			fmt.Printf("%s, %d\n\t%s\n", ld.filename, n, ld.line)
		}
	}
}

func countLines(f *os.File, counts map[lineData]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		ld := lineData{
			line:     input.Text(),
			filename: f.Name(),
		}
		counts[ld]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Problem parsing file %s\nErr: %v\n", f.Name(), err)
	}
}
