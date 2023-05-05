package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	response := "HTTP/1.1 200 OK\r\nContent-Length: 12\r\n\r\nHello world!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Server started, listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handle(conn)
	}
}
