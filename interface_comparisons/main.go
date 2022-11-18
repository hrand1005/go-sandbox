package main

import (
	"bytes"
	"fmt"
	"io"
)

// interfaceComparisons illustrates that not all interface values are comparable.
func interfaceComparisons() {
	var x interface{} = 1
	fmt.Println(x == x) // true!

	var y interface{} = []int{1, 2, 3}
	fmt.Println(y == y) // panics
}

// nilComparisons illustrates that a nil interface is different from an interface
// containing a nil value.
func nilComparisons() {
	var buf *bytes.Buffer
	fmt.Println(writerIsNil(buf)) // false!

	var w io.Writer
	fmt.Println(writerIsNil(w)) // true!
}

func writerIsNil(w io.Writer) bool {
	return w == nil
}

func main() {
	// interfaceComparisons()
	nilComparisons()
}
