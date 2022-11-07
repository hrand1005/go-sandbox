package main

import (
	"fmt"
	"strings"
)

type Rect struct {
	length int
	width  int
}

func (r *Rect) Area() int {
	return r.length * r.width
}

func (r *Rect) String() string {
	rowString := strings.Repeat("*", r.width) + "\n"
	rectString := strings.Repeat(rowString, r.length)
	return fmt.Sprintf("Length: %v, Width: %v\n\n%s", r.length, r.width, rectString)
}

func functionToMethod() {
	// create a function out of a method
	rectangleArea := (*Rect).Area

	// what would have been the receiver is now a param
	area := rectangleArea(&Rect{
		length: 5,
		width:  10,
	})

	fmt.Printf("Area: %v\n", area)
}

func interfaceVariables() {
	var s fmt.Stringer
	s = &Rect{
		length: 3,
		width:  7,
	}

	fmt.Println(s)
	// fmt.Printf("var s value:\n%v\nvar s Type:\n%T\n", s, s)
}

func main() {
	// functionToMethod()
	interfaceVariables()
}
