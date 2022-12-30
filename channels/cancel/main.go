package main

import (
	"log"
	"os"
	"time"
)

var done = make(chan struct{})

// cancelled checks whether a cancellation has occurred -- if the done channel
// is closed, it will instantly receive the zero value on that channel.
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	count(100)
}

func count(max int) {
	for i := 1; i <= max; i++ {
		if cancelled() {
			log.Fatal("gracefully shutting down after cancellation...")
		}
		log.Printf("%v\n", i)
		time.Sleep(1 * time.Second)
	}
}
