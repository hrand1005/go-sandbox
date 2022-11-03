package main

import (
	"fmt"
)

type Impl struct {
	val any
	err error
}

func (i Impl) Value() any {
	return i.val
}

func (i Impl) Error() error {
	return i.err
}

var ErrorInt = Impl{
	err: fmt.Errorf("error message"),
}

var ValidInt = Impl{
	val: 69,
}

var ValidString = Impl{
	val: "hi",
}
