package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func udp_client() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("connect server fail,err:", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请输入...")
		msg, _ := reader.ReadString('\n')

		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read reply message failed,err:", err)
			return
		}
		fmt.Println("收到回复信息，", string(reply[:n]))
	}
}
