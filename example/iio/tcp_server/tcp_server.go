package main

import (
	"fmt"
	"net"

	"github.com/mazezen/itools/iio"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		fmt.Printf("Listen tcp client err: %v", err)
		return
	}
	fmt.Println("Listen tcp client successfull ...")

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
	itcp := iio.NewITcp("")
	for {
		content, err := tcpIo.Decode(conn)
		if err != nil {
			fmt.Printf("Read from conn err: %v", err)
			break
		}
		res := string(content)
		fmt.Println(res)
	}
}
