package utils

import (
	"errors"
	"net"
)

func (s *Server) validateAndRegister(name string, connection net.Conn) error {
	//check if name is empty
	if name == "" {
		return errors.New("Name cannot be empty\n")
	}
	//check if name contains unprintables
	if !isPrintable(name) {
		return errors.New("Name cannot have unprintables\n")
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
		joinMessage := s.Clients[connection] + " has joined our chat..."
		//send history
		for _, historyMessage := range s.Messages {
			connection.Write([]byte(historyMessage + "\n"))
		}
		s.Messages = append(s.Messages, joinMessage)
		s.Mu.Unlock()
		broadcastMessage := Message{
			message:    joinMessage,
			connection: connection,
		}
		s.Channel <- broadcastMessage
		return nil
	}

	s.Mu.Unlock()

	return errors.New("Name already exists\n")
}
