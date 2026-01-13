package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"net-cat/helpers"
)

func main() {
	server := &helpers.Server{
		Broadcast: make(chan helpers.Message),
		Greeting:  helpers.ReadGreeting(),
	}
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
		portInt, err := strconv.Atoi(port)
		if err != nil || portInt <= 0 {
			fmt.Println("Invalid Port")
			return
		}
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

	go server.BroadcastMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go server.HandleClients(conn)
	}
}
