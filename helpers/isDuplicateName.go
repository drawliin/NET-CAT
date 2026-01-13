package helpers

import "strings"

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
