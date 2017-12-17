package main

import (
	"fmt"
	"net"
	"os"
)

func sender(conn net.Conn) {
	words := "get_peers123!"
	conn.Write([]byte(words))
	fmt.Println("send over")

}

func main() {
	server := "223.129.64.13:2333"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	sender(conn)

}
