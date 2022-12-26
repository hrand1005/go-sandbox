package main

import (
    "fmt"
)

func main() {
    ch := make(chan string)
    go executeQuery(ch)
    go executeQuery(ch)
    go executeQuery(ch)
    result := <-ch

    fmt.Println(result)
}

func executeQuery(ch chan<- string) {
    defer fmt.Println("exiting...")
    ch <-"result"
}
