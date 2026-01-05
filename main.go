package main

import (
	"fmt"
	"net"
	"net-cat/helpers"
	"os"
)

func main() {
	port := "8989"
	//validate args
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	} else if len(os.Args) == 2 {
		port = os.Args[1]
	}
	server := &helpers.ChatServer{
		Clients:  make(map[string]net.Conn),
		Messages: make([]string, 0),
		MsgChan:  make(chan string),
		Greeting: helpers.GetGreeting(),
	}
	server.Listen(port)

	
}
