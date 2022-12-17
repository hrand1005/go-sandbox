// Excercise 8.1
package main

import (
	"flag"
	"log"
	"net"
	"time"
)

var port = flag.String("port", "", "server port for clock service")

func main() {
	flag.Parse()

	if *port == "" {
		flag.Usage()
		return
	}

	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept conn: %v\n", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		_, err := c.Write([]byte(time.Now().Format("15:04:05\n")))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
