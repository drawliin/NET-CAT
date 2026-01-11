package utils

import "net"

func (s *Server) disconnect(connection net.Conn) {
	s.Mu.Lock()
	delete(s.Clients, connection)
	s.Mu.Unlock()
}
