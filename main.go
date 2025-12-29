package main

import (
	"bufio"
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
		go handleClients(conn)
		// conn.Write([]byte("Hello"))
	}
}

func handleClients(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			return
		}
		
		fmt.Print("Received:", msg)
		
	}
}
