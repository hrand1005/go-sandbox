package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sender(ch)
	receiver(ch)
}

func receiver(ch <-chan int) {
	for {
		if x, ok := <-ch; ok {
			fmt.Println(x)
		} else {
			fmt.Println("Channel closed, breaking...")
			break
		}
	}

	fmt.Println("Done receiving!")
}

func sender(ch chan<- int) {
	for i := 1; i < 11; i++ {
		time.Sleep(1 * time.Second)
		ch <- i
	}

	close(ch)
}
