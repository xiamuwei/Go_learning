package main

import (
	"fmt"
	"net"
)

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:3309")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
	}

	defer listener.Close()
	// 一旦有请求，立刻获取链接
	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err = ", err1)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	fmt.Println(string(buf[:n]))
}
