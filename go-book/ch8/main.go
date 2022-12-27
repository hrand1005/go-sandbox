package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	fmt.Fscanf(os.Stdin, "%v", &s)

	accum := ""
	for i := 0; i < 10; i++ {
		accum += fmt.Sprintf("%v", i)
	}
}
