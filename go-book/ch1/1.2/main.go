package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	var report strings.Builder
	for i := 1; i < len(os.Args); i++ {
		report.WriteString(fmt.Sprintf("%d: %v", i, os.Args[i]))
		if i != len(os.Args)-1 {
			report.WriteString("\n")
		}
	}
	fmt.Println(report.String())
}