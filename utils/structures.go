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
	Chan     chan string
}
