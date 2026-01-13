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

	//check if name is too long
	if len(name) >= 15 {
		return errors.New("Name should not exceed 15 characters\n")
	}

	//check if name is valid
	if !isValidName(name) {
		return errors.New("Name must only have alphanumerical characters, dash and underscore\n")
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
		//register messagee
		s.Messages = append(s.Messages, joinMessage)
		s.Mu.Unlock()
		//send to broadcast
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
