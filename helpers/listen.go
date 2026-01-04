package helpers

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var greeting = "Welcome to TCP-Chat!\n[ENTER YOUR NAME]:"

func Listen(add string) {
	//server listening
	listener, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalln("error listening on", add, ":", err)
	}else{
		fmt.Println("Listening on port :8989")
	}
	defer listener.Close()
	//handle greeting
	bytes, err := os.ReadFile("linux_logo.txt")
	if err != nil {
		fmt.Println("error reading logo")
	} else {
		greeting = string(bytes)
	}
	//create a slice for registering connections
	connections := make([]net.Conn, 0, 10)
	//create a channel for broadcasting
	ch := make(chan string)
	//create a mutex for protecting data
	var mu sync.Mutex
	//handle connections
	go BroadCastMessage(connections, &mu, ch)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("could not establish connection")
			continue
		}
		mu.Lock()
		connections = append(connections, conn)
		mu.Unlock()
		go handleConnection(conn, ch)
	}
}
