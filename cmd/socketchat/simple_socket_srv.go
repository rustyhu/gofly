package main

import (
	"fmt"
	"net"
)

// a tcp server that listens for connections and prints out the data received
// in loop.

func main() {
	const LADDR = ":8080"
	ln, err := net.Listen("tcp", LADDR)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Start listen on", LADDR)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// todo
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Received:", string(buf[:n]))

	const sendTxt = "Hello 你好 显示UTF8内容展示\n贵方客户地址 [%v].\n"
	cliAddr := conn.RemoteAddr().String()
	if _, err := conn.Write([]byte(fmt.Sprintf(
		sendTxt, cliAddr)),
	); err != nil {
		fmt.Println("Write exception -", err)
	}
}
