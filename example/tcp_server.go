package main

import (
	"fmt"
	"github.com/mazezen/itools"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		fmt.Printf("Listen tcp client err: %v", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail err: %v", err)
			continue
		}

		go read(conn)
	}
}

func read(conn net.Conn) {
	defer conn.Close()
	for {
		content, err := itools.Decode(conn)
		if err != nil {
			fmt.Printf("Read from conn err: %v", err)
			break
		}
		res := string(content)
		fmt.Println(res)
	}
}
