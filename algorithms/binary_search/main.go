package main

import (
	"fmt"
)

var orderedCollection = []int{-2, -2, 0, 1, 3, 45, 69, 70}

func main() {
	target := 69
	fmt.Printf("Searching for target %v in ordered collection:\n%+v\n", target, orderedCollection)
	result := int(BinarySearch(target, orderedCollection))
	if result != NotFound {
		fmt.Printf("Target found at index '%v'.\n", result)
	} else {
		fmt.Printf("Target '%v' not found.\n", target)
	}
}
