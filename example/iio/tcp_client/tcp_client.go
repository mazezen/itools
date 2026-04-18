package main

import (
	"fmt"
	"net"
	"time"

	"github.com/mazezen/itools/iio"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	defer conn.Close()
	if err != nil {
		fmt.Printf("Connect tcp err: %v", err)
		return
	}

	itcp := iio.NewITcp("")
	for {
		err = itcp.Encode(conn, "Hi mary!!!")
		time.Sleep(time.Second * 5)
		if err != nil {
			fmt.Printf("Unpack err: %v", err)
			return
		}
	}
}
