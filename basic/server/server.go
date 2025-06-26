package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	logErr = log.New(os.Stdout, "Error: ", log.Ldate|log.Ltime)
)

const (
	HOSTNAME  = "localhost"
	PORT      = "8000"
	NETWORK_T = "tcp4"
)

const (
	MAX_REQ_SIZE = 100
	MAX_RES_SIZE = 100
)

func exitIfErr(err error) {
	if err != nil {
		logErr.Println(err.Error())
		os.Exit(1)
	}
}

func main() {
	listener, err := net.Listen(NETWORK_T, HOSTNAME+":"+PORT)
	exitIfErr(err)
	defer listener.Close()

	for {
		logger.Println("Waiting for connection in", listener.Addr().String())

		conn, err := listener.Accept()
		exitIfErr(err)

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var (
		b   bytes.Buffer
		req = make([]byte, MAX_REQ_SIZE)
		res = make([]byte, MAX_RES_SIZE)
	)

	_, err := conn.Read([]byte(req))
	exitIfErr(err)

	logger.Println("Request recv: ", string(req))

	fmt.Fprintf(&b, "Hello")
	res = b.Bytes()

	_, err = conn.Write([]byte(res))
	exitIfErr(err)

	logger.Println("Closing connection: ", conn.LocalAddr().String())
}
