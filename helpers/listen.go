package helpers

import (
	"fmt"
	"log"
	"net"
)

func (s *ChatServer) Listen(port string) {
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print("Listening on the port :", port, "\n")
	//ready for broadcasting
	go s.broadCastMessages()
	//accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handleConnection(conn)
	}
}
