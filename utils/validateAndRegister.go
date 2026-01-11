package utils

import (
	"errors"
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
	nameExists := false

	for _, existingName := range s.Clients {
		if existingName == name {
			nameExists = true
			break
		}
	}

	if !nameExists {
		s.Clients[connection] = name
		s.Mu.Unlock()
		return nil
	}

	s.Mu.Unlock()

	return errors.New("Name already exists\n")
}
