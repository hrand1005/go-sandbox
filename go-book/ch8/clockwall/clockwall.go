// Exercise 8.1
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type clock struct {
	name   string
	server string
}

func parseClocks(args []string) []clock {
	clocks := make([]clock, 0, len(args))
	for _, v := range args {
		tuple := strings.Split(v, "=")
		clocks = append(clocks, clock{
			name:   tuple[0],
			server: tuple[1],
		})
	}

	return clocks
}

type clockWall struct {
	header string
	conns  []net.Conn
}

func (cw *clockWall) Add(c clock) error {
	conn, err := net.Dial("tcp", c.server)
	if err != nil {
		return err
	}
	cw.header += c.name + "\t\t"
	cw.conns = append(cw.conns, conn)

	return nil
}

func (cw *clockWall) Header() string {
	return cw.header
}

func (cw *clockWall) Current() string {
	times := ""
	for _, conn := range cw.conns {
		var t string
		fmt.Fscanf(conn, "%v", &t)

		times += t + "\t"
	}

	return fmt.Sprintf("\r%v", times)
}

func (cw *clockWall) Close() {
	for _, conn := range cw.conns {
		conn.Close()
	}
}

func main() {
	clocks := parseClocks(os.Args[1:])

	cw := &clockWall{}
	defer cw.Close()

	for _, c := range clocks {
		if err := cw.Add(c); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(cw.Header())
	for {
		fmt.Printf("%s", cw.Current())
	}
}
