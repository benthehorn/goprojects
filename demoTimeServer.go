package main 

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	
	//listen to a port
	ln, err := net.Listen("tcp", ":8080")
	defer ln.Close()
	if err != nil {
		log.Panic(err)
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}

		go func(conn net.Conn) {
			daytime := time.Now().String() + " ... Hello\n"

			fmt.Fprintf(conn, daytime)
			conn.Close()
		}(c)
	}
}