package main

import (
	"fmt"
	"net"
)

func main() {
	// 拨号
	conn, err := net.Dial("tcp", "127.0.0.1:3309")
	if err != nil {
		fmt.Println("net.Dial err =", err)
	}

	defer conn.Close()

	conn.Write([]byte("hello golang"))
}
