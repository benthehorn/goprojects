package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8080")
    defer conn.Close()
    if err != nil {
        log.Panic(err)
    }

    connbuf := bufio.NewReader(conn)
    for {
        str, err := connbuf.ReadString('\n')
        if len(str) > 0 {
            fmt.Println(str)
            break
        }
        if err != nil {
            break
        }
    }
}