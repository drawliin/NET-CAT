package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Client struct {
	c_conn net.Conn
	c_name string
}

var (
	clients []Client
	mu      sync.Mutex
)

var broadcast = make(chan string)

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

	go func() {
		for msg := range broadcast {
			mu.Lock()
			for _, c := range clients {
				c.c_conn.Write([]byte(msg))
			}
			mu.Unlock()
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClients(conn)
	}
}

func handleClients(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	conn.Write([]byte("Welcome To The TCP Chat !!!\n"))
	conn.Write([]byte("Enter Your Name: "))

	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if err != nil {
		fmt.Println("Client disconnected:", conn.RemoteAddr())
		return
	}
	if name == "" {
		conn.Write([]byte("Name Cannot be empty"))
		return
	}
	conn.Write([]byte("You name " + name))
	// store client infos
	client := Client{
		c_conn: conn,
		c_name: name,
	}
	mu.Lock()
	if len(clients) >= 10 {
		mu.Unlock()
		conn.Write([]byte("Chat is full\n"))
		return
	}
	clients = append(clients, client)
	mu.Unlock()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error when reading messages")
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formatted := fmt.Sprintf("[%s][%s]:%s\n", timestamp, client.c_name, msg)
		broadcast <- formatted
	}
}
