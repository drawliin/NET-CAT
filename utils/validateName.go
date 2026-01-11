package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func (s *Server) validateAndRegister(name string, connection net.Conn) error {
	//trim whitespace
	name = strings.TrimSpace(name)

	//check if name is empty
	if name == "" {
		return errors.New("Name cannot be empty\n")
	}

	//check if name already exists
	s.Mu.Lock()

	_, exists := s.Clients[name]

	if !exists {
		s.Clients[name] = connection
		s.Mu.Unlock()
		return nil
	}

	s.Mu.Unlock()

	return errors.New("Name already exists\n")
}
