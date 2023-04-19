package main

import (
	"testing"
)

var stringArgs = []string{
	"hi", "there", "how", "are", "you", "I",
	"hope", "you", "have", "a", "nice", "day",
}

func BenchmarkConcatArgs1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatArgs1(stringArgs)
	}
}

func BenchmarkConcatArgs2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatArgs2(stringArgs)
	}
}

func BenchmarkConcatArgs3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatArgs3(stringArgs)
	}
}
