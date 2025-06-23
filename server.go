package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.0:8000")
	if err != nil {
		fmt.Println("net.Listen")
		return
	}
	defer listener.Close()

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("listner.Accept")
		return
	}
	defer connection.Close()

}
