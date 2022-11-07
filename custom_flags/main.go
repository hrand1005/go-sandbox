package main

import (
  "flag"
  "fmt"
)

var rect = RectFlag("rect", Rect{}, "rectangle as defined by LxW")

func main() {
  flag.Parse()
  fmt.Println(*rect)
}
