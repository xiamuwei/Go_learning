package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listen.Close()

	buf := make([]byte, 1024)
	num, addr, err1 := listen.ReadFromUDP(buf)
	if err1 != nil {
		fmt.Println("listen.ReadFromUDP err = ", err1)
		return
	}
	fmt.Printf("data:%v addr:%v count:%v\n", string(buf[:num]), addr, num)

}
