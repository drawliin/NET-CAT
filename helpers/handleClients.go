package helpers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
	"unicode/utf8"
)

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
		if len(name) > 15 {
			if _, err = conn.Write([]byte("Name should not exceed 15 characters\n")); err != nil {
				return
			}
			continue
		}
		if !isValidName(name) {
			if _, err = conn.Write([]byte("Name must only have alphanumerical characters, dash and underscore\n")); err != nil {
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
	client := Client{
		c_conn: conn,
		c_name: name,
		send:   make(chan string, 16),
	}
	s.mu.Lock()

	if len(s.clients) >= 10 {
		s.mu.Unlock()
		conn.Write([]byte("Chat is full\n"))
		return
	}
	s.clients = append(s.clients, &client)
	s.mu.Unlock()

	go client.clientWriter()

	// load old messages to new clients
	s.mu.Lock()
	for _, oldMssg := range s.history {
		client.send <- oldMssg
	}
	s.mu.Unlock()

	userJoinMssg := fmt.Sprintf("%s has joined our chat...\n", client.c_name)

	s.mu.Lock()
	for _, c := range s.clients {
		if c != &client {
			select {
			case c.send <- userJoinMssg:
			default:
				c.c_conn.Close()
			}
		}
	}
	s.mu.Unlock()

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			s.removeClient(conn)
			leaveMsg := fmt.Sprintf("%s has left our chat...\n", client.c_name)
			s.Broadcast <- leaveMsg
			close(client.send)
			break
		}
		msg = strings.TrimSpace(msg)
		if msg == "" || !isPrintable(msg) {
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
