package utils

import (
	"fmt"
	"time"
)

func (s *Server) BroadcastMessage() {
	for message := range s.Channel {
		s.Mu.Lock()
		time := time.Now().Format("2006-01-02 15:04:05")
		for connection := range s.Clients {
			//check if it's sender
			if message.connection == connection {
				prompt := fmt.Sprintf("[%s][%s]: ", time, s.Clients[message.connection])
				connection.Write([]byte(prompt))
				continue
			}
			connection.Write([]byte("\n" + message.message))
			//if chat state is join or leave
			if message.chatState {
				prompt := fmt.Sprintf("[%s][%s]: ", time, s.Clients[connection])
				connection.Write([]byte("\n" + prompt))
			}
		}
		s.Mu.Unlock()
	}
}
