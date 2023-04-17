package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]

	concatTimer("printArgs1", printArgs1, args)
	fmt.Println()
	concatTimer("printArgs2", printArgs2, args)
	fmt.Println()
	concatTimer("printArgs3", printArgs3, args)
}

// concatTimer prints the execution time and result of the provided 
// concatenation function. Expects name, function, and function args.
func concatTimer(name string, f func([]string) string, fargs []string) {
	start := time.Now()
	s := f(fargs)
	duration := time.Since(start)

	fmt.Printf("%s\nresult: %s\nexecution time: %v\n", name, s, duration)
}

// printArgs1 builds a string form the provided arguments through addition. 
// I decided to return a string because I don't think printing to stdout is 
// particularly useful during profiling. 
func printArgs1(args []string) string {
	s, sep := "", ""
	for _, a := range args {
		s += sep + a
		sep = ", "
	}
	return s
}

// printArgs2 is like printArgs1 except it uses strings.Builder.
func printArgs2(args []string) string {
	var s strings.Builder
	var sep string
	for _, a := range args {
		s.WriteString(sep + a)
		sep = ", "
	}
	return s.String()
}

// printArgs3 simply uses strings.Join.
func printArgs3(args []string) string {
	return strings.Join(args, ", ")
}