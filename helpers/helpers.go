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
	history []string
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
	if err != nil {
		fmt.Println("can't establish connection:", conn.RemoteAddr())
		return
	}
	name = strings.TrimSpace(name)
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

	// load old messages to new clients
	s.mu.Lock()
	for _, oldMssg := range s.history {
		conn.Write([]byte(oldMssg))
	}
	s.mu.Unlock()
	

	userJoinMssg := fmt.Sprintf("%s has joined the group\n", client.c_name)
	s.Broadcast <- userJoinMssg


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
		s.mu.Lock()
		s.history = append(s.history, formatted)
		s.mu.Unlock()
		s.Broadcast <- formatted
	}
}

// BroadCast a Message to all clients
func (s *Server) BroadcastMessages() {
	
	for msg := range s.Broadcast {
		// create copy of clients to not hold the mutex so long
		s.mu.Lock()
		clientsCopy := append([]Client(nil), s.clients...)
		s.mu.Unlock()
		for _, c := range clientsCopy {
			c.c_conn.Write([]byte(msg))
		}
	}
}

// Remove a client if disconnected
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
