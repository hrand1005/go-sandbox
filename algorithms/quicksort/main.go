package main

import (
	"fmt"
)

// sample quicksort inputs
var intSlice = []int{2, 4, 5, 1, 0, -1, 10}
var stringSlice = []string{"b", "hi", "herb", "zebra", "mom", "apple"}

func main() {
	sortElements(intSlice)
	sortElements(stringSlice)
}

// sortElements prints slice before and after sorting
func sortElements[T sortable](slice []T) {
	fmt.Printf("Sorting slice with elements of type %T:\n\n", slice[0])
	fmt.Printf("Before:\n\t%v\n\n", slice)

	quicksort(slice, 0, len(slice)-1)

	sep := "------------\n\n"
	fmt.Printf("After:\n\t%v\n\n%s", slice, sep)
}
