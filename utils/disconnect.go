package utils

import (
	"log"
	"net"
)

func (s *Server) disconnect(connection net.Conn) {
	name := s.Clients[connection]
	s.Mu.Lock()
	delete(s.Clients, connection)
	leaveMessage := name + " has left our chat...\n"
	log.Println("client disconnected")
	s.Mu.Unlock()
	s.Channel <- leaveMessage
}
