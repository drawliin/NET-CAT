package helpers

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

)

type Message struct {
	client  *Client
	message string
}

type Server struct {
	clients   []*Client
	mu        sync.Mutex
	Broadcast chan Message
	history   []string
	Greeting  []byte
}

type Client struct {
	c_conn net.Conn
	c_name string
	send   chan struct{
		message string
		history bool
	}
}

func (s *Server) HandleClients(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	if _, err := conn.Write(s.Greeting); err != nil {
		return
	}

	var name string
	var err error
	for {
		if _, err = conn.Write([]byte("[ENTER YOUR NAME]: ")); err != nil {
			return
		}
		name, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("can't establish connection:", conn.RemoteAddr())
			return
		}
		name = strings.TrimSpace(name)
		if name == "" {
			if _, err = conn.Write([]byte("Name Cannot be empty\n")); err != nil {
				return
			}
			continue
		}
		if s.isDuplicateName(name) {
			if _, err = conn.Write([]byte("Name is already taken\n")); err != nil {
				return
			}
			continue
		}
		if utf8.RuneCountInString(name) > 30 {
			if _, err = conn.Write([]byte("Name is too long\n")); err != nil {
				return
			}
			continue
		}
		break
	}

	// store client infos
	client := &Client{
		c_conn: conn,
		c_name: name,
		send:   make(chan struct{
			message string
			history bool
		}, 16),
	}
	s.mu.Lock()

	if len(s.clients) >= 10 {
		s.mu.Unlock()
		conn.Write([]byte("Chat is full\n"))
		return
	}
	s.clients = append(s.clients, client)
	s.mu.Unlock()

	go client.clientWriter()

	// load old messages to new clients
	s.mu.Lock()
	for _, oldMssg := range s.history {
		client.send <- struct{message string; history bool}{message: oldMssg, history: true}
	}
	s.mu.Unlock()

	userJoinMssg := fmt.Sprintf("%s has joined our chat...\n", client.c_name)
	s.Broadcast <- Message{client: client, message: userJoinMssg}
	// s.mu.Lock()
	// for _, c := range s.clients {
	// 	if c != &client {
	// 		select {
	// 		case c.send <- userJoinMssg:
	// 		default:
	// 			c.c_conn.Close()
	// 		}
	// 	}
	// }
	// s.mu.Unlock()

	for {
		promptTime := time.Now().Format("2006-01-02 15:04:05")
		prompt := fmt.Sprintf("[%s][%s]:", promptTime, client.c_name)
		if _, err = conn.Write([]byte(prompt)); err != nil {
			return
		}
		msg, err := reader.ReadString('\n')
		if err != nil {
			s.removeClient(conn)
			leaveMsg := fmt.Sprintf("%s has left our chat...\n", client.c_name)
			s.Broadcast <- Message{client: client, message: leaveMsg}
			close(client.send)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" {
			continue
		}
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		formatted := fmt.Sprintf("\n[%s][%s]:%s\n", timestamp, client.c_name, msg)
		s.mu.Lock()
		s.history = append(s.history, formatted)
		s.mu.Unlock()
		s.Broadcast <- Message{client: client, message: formatted}
	}
}

// BroadCast a Message to all clients
func (s *Server) BroadcastMessages() {
	for msg := range s.Broadcast {
		// create copy of clients to not hold the mutex so long
		s.mu.Lock()
		clientsCopy := append([]*Client(nil), s.clients...)
		s.mu.Unlock()

		for _, c := range clientsCopy {
			if c == msg.client {
				continue
			}
			select {
			case c.send <- struct{message string; history bool}{message: msg.message, history: false}:
				// Success: Message added to buffer
			default:
				// Failure: Buffer is full (16/16).
				// Client is lagging.
				fmt.Println("Kicking lagging client:", c.c_name)
				c.c_conn.Close()
			}
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

func (c *Client) clientWriter() {
	for msg := range c.send {
		var prompt string
		if !msg.history {
			promptTime := time.Now().Format("2006-01-02 15:04:05")
			prompt = fmt.Sprintf("[%s][%s]:", promptTime, c.c_name)
		}
		_, err := c.c_conn.Write([]byte("\n" + msg.message + prompt))
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}
	}
}

// check if name already exists
func (s *Server) isDuplicateName(name string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, client := range s.clients {
		if strings.EqualFold(client.c_name, name) {
			return true
		}
	}
	return false
}

func ReadGreeting() []byte {
	defaultGreeting := "Welcome to TCP-Chat!\n[ENTER YOUR NAME]:"
	logoBytes, err := os.ReadFile("greeting.txt")
	if err != nil {
		return []byte(defaultGreeting)
	}
	return logoBytes
}
