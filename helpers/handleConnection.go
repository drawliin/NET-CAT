package helpers

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func (s *ChatServer) handleConnection(conn net.Conn) {
	defer conn.Close()
	writer := bufio.NewWriter(conn)
	reader := bufio.NewReader(conn)
	name := ""
	invalidName := true
	for invalidName {
		//greet client
		writer.WriteString(s.Greeting)
		writer.Flush()
		//read client name
		var err error
		name, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		name = strings.TrimSpace(name)
		//check if name is empty
		if name == "" {
			continue
		}
		//check if name already exists
		s.Mu.Lock()
		_, exists := s.Clients[name]
		s.Mu.Unlock()
		if !exists {
			invalidName = false
		}
	}
	//add the conn to the map
	s.Mu.Lock()
	s.Clients[name] = conn
	s.Mu.Unlock()
	joinMessage := fmt.Sprintln(name, "has joined our chat...")
	s.MsgChan <- joinMessage
	defer func() {
		s.Mu.Lock()
		delete(s.Clients, name)
		s.Mu.Unlock()
		leaveMessage := fmt.Sprintln(name, "has left our chat...")
		s.MsgChan <- leaveMessage
	}()
	//handle messages
	for {
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message)
		if err != nil {
			break
		}
		if message == "" {
			continue
		}
		s.MsgChan <- message
	}
}
