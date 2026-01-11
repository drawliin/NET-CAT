package main

import (
	"log"
	"net"
	"net-cat/utils"
	"os"
)

func main() {
	//check if number of arguments is higher than one
	if len(os.Args) > 2 {
		log.Println("[USAGE]: ./TCPChat $port")
		return
	}

	port := "8989"

	//collect port from arguments if it exists
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	//create an instance of server
	server := utils.Server{
		Chan:     make(chan string),
		Clients:  make(map[net.Conn]string),
		Greeting: utils.Greeting(),
	}

	//start listening
	server.Listen(port)
}
