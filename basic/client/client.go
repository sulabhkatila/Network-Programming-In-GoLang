package main

import (
	"fmt"
	"net"
	"strconv"
)

var PORT = 8000

func main() {
	conn, err := net.Dial("tcp", "127.0.0.0:"+strconv.Itoa(PORT))
	if err != nil {
		fmt.Println("failed to connect")
		return
	}

	fmt.Println("Connected")
	go handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	req := "Hello"
	_, err := conn.Write([]byte(req))
	if err != nil {
		fmt.Println("Failed to write")
		conn.Close()
		return
	}
	fmt.Println("Wrote: ", req)

	var res string
	_, error := conn.Read([]byte(res))
	if error != nil {
		fmt.Println("Failed to read")
		conn.Close()
		return
	}

	fmt.Println("read: ", res)
	conn.Close()
}
