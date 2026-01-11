package utils

import (
	"bufio"
	"log"
	"net"
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

	for {
		_, err := connection.Write([]byte(askNameMessage))

		if err != nil {
			log.Printf("failed to register client...\n%s\n", err)
			return
		}

		name, err := reader.ReadString('\n')

		if err != nil {
			log.Printf("failed to register client...\n%s\n", err)
			return
		}

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
		log.Println("client registered:", name)
		break
	}

}
