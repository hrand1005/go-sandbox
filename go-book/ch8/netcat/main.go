package main

import (
    "io"
    "log"
    "net"
    "os"
)

// netcat3 -- used in 8.4.1, used in exercise 8.3
func main() {
    conn, err := net.Dial("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    done := make(chan struct{})
    go func() {
        io.Copy(os.Stdout, conn)
        log.Println("done")
        done <-struct{}{}
    }()
    mustCopy(conn, os.Stdin)
    conn.Close()
    <-done
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        log.Fatal(err)
    }
}
