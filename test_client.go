package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"./protocol"
)

func sender(conn net.Conn) {
	buffer := make([]byte, 1024)
	for i := 0; i < 1000; i++ {
		words := "{\"Id\":i,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(protocol.Packet([]byte(words)))
		conn.Read(buffer)
		// println(strings.Replace(string(buffer), " ", "", -1))
		println(string(buffer)[:20])
	}

	runtime.Stack(buffer, false)
	fmt.Println(string(buffer)[:10], "send over")
}

func main() {
	server := "127.0.0.1:9988"
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

	defer conn.Close()
	fmt.Println("connect success")
	go sender(conn)
	go sender(conn)
	for {
		time.Sleep(1 * 1e9)
	}
}
