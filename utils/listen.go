package utils

import (
	"log"
	"net"
)

func (s *Server) Listen(port string) {
	//start listening on the given address
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("error listening on port %s...\n%s", port, err)
	}
	log.Println("listening on port", port)

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
