package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sampleSlices := buildSampleSlices()

	start := time.Now()
	synchronousSum(sampleSlices)
	elapsed := time.Since(start)

	fmt.Printf("synchronousSum elapsed in %v\n", elapsed)

	start2 := time.Now()
	concurrentSum(sampleSlices)
	elapsed2 := time.Since(start2)

	fmt.Printf("concurrentSum elapsed in %v\n", elapsed2)
}

func synchronousSum(allSlices [][]int) {
	sum := 0
	for _, s := range allSlices {
		for _, v := range s {
			sum += v
		}
	}

	fmt.Println(sum)
}

func concurrentSum(allSlices [][]int) {
	in := make(chan []int)
	out := make(chan int)
	go sum(in, out)

	for _, s := range allSlices {
		in <- s
	}
	close(in)

	fmt.Println(<-out)
}

// buildSampleSlices creates 10 slices, each with 100 elements. The sum of all
// elements in all slices is _
func buildSampleSlices() [][]int {
	full := make([][]int, 10)
	for i := 0; i < 10; i++ {
		inner := make([]int, 100)
		for j := 0; j < 100; j++ {
			inner[j] = j
		}
		full[i] = inner
	}

	return full
}

// sum adds all elements of all slices pushed to the provided int-slice channel.
func sum(sliceChan <-chan []int, result chan<- int) {
	wg := sync.WaitGroup{}

	ch := make(chan int)
	for s := range sliceChan {
		wg.Add(1)
		go func(s []int) {
			defer wg.Done()
			summer(s, ch)
		}(s)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	grandTotal := 0
	for subTotal := range ch {
		grandTotal += subTotal
	}

	result <- grandTotal
}

// summer sums all the elements in one slice
func summer(ints []int, total chan<- int) {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	total <- sum
}
