package main

import (
	"fmt"
	"os"
)

const maxCount = 25

func main() {
	counter := make(chan int)
	fib := make(chan int)

	go countStage(maxCount, counter)
	go fibStage(fib, counter)
	printStage(fib)
}

// countStage counts up to 'max' and pushes the results to the 'out' channel.
// Closes the channel after completing.
func countStage(max int, out chan<- int) {
	for i := 0; i < max+1; i++ {
		out <- i
	}
	close(out)
}

// fibStage calculates the nth term in the fibonacci sequence for each term
// received on 'in' chan, and pushes the result to 'out' chan.
func fibStage(out chan<- int, in <-chan int) {
	for n := range in {
		out <- fibN(n)
	}
	close(out)
}

// fibN returns the nth numberin the fibonacci sequence.
func fibN(n int) int {
	if n <= 1 {
		return 1
	}
	return fibN(n-2) + fibN(n-1)
}

// printStage reads integers from 'in' and prints them to stdout.
func printStage(in <-chan int) {
	for res := range in {
		fmt.Fprintln(os.Stdout, res)
	}
}
