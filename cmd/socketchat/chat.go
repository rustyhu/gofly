package main

import (
	"fmt"
	"io"
	"net"
)

const (
	BufLen = 2048
)

type Mode uint8

const (
	DIAL Mode = iota
	WAIT
	LAST // need a config file to store last choice
	QUIT
)

type BasicConf struct {
	mode   Mode
	thisIP string
	peerIP string
	// 接收端口
	ListenPort int
	SendPort   int
}

func connectToOthers(b *BasicConf) {
	// TCP client mode
	fmt.Println("Net address you want to connect:")
	fmt.Scanln(b.peerIP)
}

func waitForConnects(b *BasicConf) {
	// TCP server mode
	listenner, _ := net.Listen("tcp", fmt.Sprintf("%s:%d", b.thisIP, b.ListenPort))
	fmt.Printf("Start listen on local %s on port %d\n", b.thisIP, b.ListenPort)

	buf := [BufLen]byte{}
	for {
		c, _ := listenner.Accept()
		fmt.Println("Get remote connection from:", c.RemoteAddr().String())

		for {

			// arrange 2 status: recv and send
			for {
				n, err := c.Read(buf[:])
				if n > 0 && err == nil {
					fmt.Println("Recv:", string(buf[:]))
				} else if err == io.EOF {
					break
				} else {
					fmt.Println("Recv error:", err)
					c.Close()
				}
			}
		}
	}
}

func main() {
	// only implement file transferation
	var basicConf BasicConf

OUTER:
	for {
		fmt.Println(
			`Welcome to chatting room!
Your choice:
1. Connect to others
2. Wait for others to connect
3. Do tha last option again
4. Quit
Input: `)

		var choice uint
		fmt.Scanln(choice)

		switch choice {
		case uint(DIAL):
			basicConf.mode = DIAL
			connectToOthers(&basicConf)
		case uint(WAIT):
			basicConf.mode = WAIT
			waitForConnects(&basicConf)
		case uint(QUIT):
			break OUTER
		default:
			fmt.Println("Invalid option")
		}
	}
}
