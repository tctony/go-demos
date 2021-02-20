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
	listenAddr := ":0"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", listenAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	exitOnError("listen", err)
	fmt.Printf("listening on %s\n", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	nr, err := conn.Read(buf)
	if err == nil && nr > 0 {
		fmt.Printf("message from client: %q\n", buf[:nr])
		conn.Write([]byte(fmt.Sprintf("echo of %q", buf[:nr])))
	}
}
