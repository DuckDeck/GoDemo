package main

import (
	"fmt"
	"net"
	"strings"
)

func udp_server() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("Listen UDP failed,err:", err)
		return
	}
	defer conn.Close()
	//不需要建立连接，直接 收发数据
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read from udp failed,err", err)
			return
		}
		fmt.Println(data[:n])
		replay := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(replay), addr)
	}
}
