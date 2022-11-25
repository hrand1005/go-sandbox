package main

import (
	"fmt"
)

var intSlice = []int{1, -1, 3, 4, 1, 0}

func main() {
	fmt.Printf("Before sorting:\n%+v\n", intSlice)
	sortedSlice := mergesort(intSlice)
	fmt.Printf("After sorting:\n%+v\n", sortedSlice)
}
