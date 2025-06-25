package main

import (
	"fmt"
	"net"
	"strconv"
)

var PORT = 8000

func main() {
	listener, err := net.Listen("tcp", "127.0.0.0:"+strconv.Itoa(PORT))
	if err != nil {
		fmt.Println("net.Listen")
		return
	}
	defer listener.Close()

	for {
		fmt.Println("Waiting for connection")
		conn, err := listener.Accept()
		fmt.Println("Connection Accepted: ", conn.LocalAddr().String())
		if err != nil {
			fmt.Println("listner.Accept")
			return
		}

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var req string
	_, err := conn.Read([]byte(req))
	if err != nil {
		conn.Close()
		fmt.Println("Closing connection: ", conn.LocalAddr().String())
		return
	}

	fmt.Println("Request recv: ", req)
	conn.Close()
	fmt.Println("Closing connection: ", conn.LocalAddr().String())
}
