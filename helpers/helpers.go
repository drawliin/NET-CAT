package helpers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	clients   []Client
	mu        sync.Mutex
	Broadcast chan string
}

type Client struct {
	c_conn net.Conn
	c_name string
}

func (s *Server) HandleClients(conn net.Conn) {
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
	conn.Write([]byte("Your name " + name + "\n"))
	// store client infos
	client := Client{
		c_conn: conn,
		c_name: name,
	}
	s.mu.Lock()
	if len(s.clients) >= 10 {
		s.mu.Unlock()
		conn.Write([]byte("Chat is full\n"))
		return
	}
	s.clients = append(s.clients, client)
	s.mu.Unlock()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			s.removeClient(conn)
			leaveMsg := fmt.Sprintf("%s has left the chat...\n", client.c_name)
			s.Broadcast <- leaveMsg
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formatted := fmt.Sprintf("[%s][%s]:%s\n", timestamp, client.c_name, msg)
		s.Broadcast <- formatted
	}
}

func (s *Server) BroadcastMessages() {
	for msg := range s.Broadcast {
		s.mu.Lock()
		for _, c := range s.clients {
			c.c_conn.Write([]byte(msg))
			fmt.Println(msg)
		}
		s.mu.Unlock()
	}
}

func (s *Server) removeClient(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, c := range s.clients {
		if c.c_conn == conn {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			return
		}
	}
}
