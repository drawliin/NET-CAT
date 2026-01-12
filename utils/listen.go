package utils

import (
	"fmt"
	"log"
	"net"
)

func (s *Server) Listen(port string) {
	//start listening on the given address
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("error listening on port %s...\n%s", port, err)
	}
	fmt.Printf("Listening on the port :%s\n", port)
	//broadcast messages
	go s.BroadcastMessage()
	//accept connections
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Printf("could not accept connection...\n%s\n", port)
		}

		//handle connection
		go s.handleConnection(connection)
	}

}
