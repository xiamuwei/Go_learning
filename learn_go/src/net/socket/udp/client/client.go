package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("net.DialUDP err = ", err)
		return
	}
	_, err = socket.Write([]byte("hello udp protocl"))
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return
	}
}
