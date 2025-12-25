package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on the port :" + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("New connection from", conn.RemoteAddr())
		conn.Close()
	}
}
