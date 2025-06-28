package main

import (
	"fmt"
	"net"
	"os"
)

const (
	NETWORKTYPE = "tcp4"
	HOSTADDRESS = "localhost"
	PORT        = "8000"
)

func main() {
	conn, err := net.Dial(NETWORKTYPE, HOSTADDRESS+":"+PORT)
	exitIfError(err)

	fmt.Println("Connected")
	go handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	req := "Hello"
	_, err := conn.Write([]byte(req))
	exitIfError(err)

	fmt.Println("Wrote: ", req)

	var res string
	_, err = conn.Read([]byte(res))
	exitIfError(err)

	fmt.Println("read: ", res)
}

func exitIfError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
