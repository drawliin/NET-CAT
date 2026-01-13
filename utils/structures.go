package utils

import (
	"net"
	"sync"
)

type Server struct {
	Clients  map[net.Conn]string
	Messages []string
	Greeting string
	Mu       sync.Mutex
	Channel  chan Message
}

type Message struct {
	message    string
	connection net.Conn
}
