package utils

import (
	"log"
	"net"
)

func (s *Server) disconnect(connection net.Conn) {
	s.Mu.Lock()
	delete(s.Clients, connection)
	log.Println("client disconnected")
	s.Mu.Unlock()
}
