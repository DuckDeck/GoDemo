package main

import (
	"fmt"
	"net"
)

func tcp_server() {

	lisntener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("Start tcp server on 127.0.0.1:20000 fail,error", err)
		return
	}

	for { //这样就能有多个接收者
		conn, err := lisntener.Accept()
		fmt.Println("accept listenter")
		if err != nil {
			fmt.Println("accept fail,error", err)
			return
		}
		go processConn(conn)
	}
}

func processConn(conn net.Conn) {
	var tmp [128]byte
	for { //这样就可以接收多个消息
		n, err := conn.Read(tmp[:])
		fmt.Println("get a message")
		if err != nil {
			fmt.Println("read from conn failed", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
