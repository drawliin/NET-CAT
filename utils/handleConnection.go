package utils

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func (s *Server) handleConnection(connection net.Conn) {
	//greet and ask the client for a name
	_, err := connection.Write([]byte(s.Greeting + "\n"))
	//check if there is space
	s.Mu.Lock()
	if len(s.Clients) >= 10 {
		s.Mu.Unlock()
		log.Println("chat is full...")
		connection.Write([]byte("Chat is full, please try again later...\n"))
		return
	}
	s.Mu.Unlock()
	if err != nil {
		log.Printf("failed to register client...\n%s\n", err)
		return
	}

	//read the given name until it's valid
	askNameMessage := "[ENTER YOUR NAME]: "
	reader := bufio.NewReader(connection)
	name := ""
	for {
		_, err := connection.Write([]byte(askNameMessage))

		if err != nil {
			log.Printf("failed to register client...\n%s\n", err)
			return
		}

		name, err = reader.ReadString('\n')

		if err != nil {
			log.Printf("failed to register client...\n%s\n", err)
			return
		}

		//trim whitespace
		name = strings.TrimSpace(name)

		//validate name
		err = s.validateAndRegister(name, connection)

		//show error
		if err != nil {
			_, err = connection.Write([]byte(err.Error()))

			if err != nil {
				log.Printf("failed to register client...\n%s\n", err)
				return
			}

			continue
		}
		break
	}

	//listen for messages on the connection
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			s.disconnect(connection)
			return
		}

		//broadcast message
		s.Mu.Lock()
		time := time.Now().Format("2006-01-02 15:04:05")
		formattedMessage := fmt.Sprintf("[%s][%s]:%s", time, name, message)
		s.Messages = append(s.Messages, formattedMessage)
		s.Mu.Unlock()
		broadcastMessage := Message{
			message:    formattedMessage,
			connection: connection,
		}
		s.Channel <- broadcastMessage
	}

}
