package main

import (
	"fmt"
	"net"
	"os"
)

func exitOnError(info string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s error: %s", info, err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", serverAddr)
	exitOnError("resolve addr", err)

	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	exitOnError("tcp conn", err)

	_, err = tcpConn.Write([]byte("hello"))
	exitOnError("write", err)

	resp := make([]byte, 1024)
	nr, err := tcpConn.Read(resp)
	exitOnError("read", err)

	fmt.Printf("receive %d bytes from server: %q", nr, resp[:nr])

	os.Exit(0)
}
