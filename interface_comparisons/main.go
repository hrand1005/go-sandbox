package main

import (
  "fmt"
)

func main() {
  var x interface{} = 1
  fmt.Println(x == x) // true!

  var y interface{} = []int{1, 2, 3}
  fmt.Println(y == y) // panics
}
