package utils

func (s *Server) BroadcastMessage() {
	for message := range s.Channel {
		s.Mu.Lock()
		for connection := range s.Clients {
			connection.Write([]byte(message))
		}
		s.Mu.Unlock()
	}
}
