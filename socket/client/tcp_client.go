package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func tcp_server() {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dail 127.0.0.1:20000 failed,error", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var msg string
	for {
		fmt.Println("请说话 ：")
		text, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(text)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}
