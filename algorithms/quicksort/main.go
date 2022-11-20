package main

import (
	"fmt"
)

// sample quicksort input
var intSlice = []int{2, 4, 5, 1, 0, -1, 10}

// var intSlice = []int{1, 4, 2}

func main() {
	fmt.Printf("int slice before sorting:\n%#v\n", intSlice)
	quicksort(intSlice, 0, len(intSlice)-1)
	fmt.Printf("int slice after sorting:\n%#v\n", intSlice)
}
