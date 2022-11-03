package main

import (
	"fmt"
)

type Result interface {
	Value() any
	Error() error
}

func Do(f ...func() Result) Result {
	var last Result
	for _, v := range f {
		if last = v(); last.Error() != nil {
			return last
		}
	}
	return last
}

func main() {

	res := Do(
		func() Result { return ValidInt },
		func() Result { return ErrorInt },
		func() Result { return ValidString },
	)

	fmt.Printf("Value: %v\nError: %s\n", res.Value(), res.Error())
}
