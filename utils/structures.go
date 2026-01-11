package utils

import (
	"net"
	"sync"
)

type Server struct {
	Clients  map[string]net.Conn
	Messages []string
	Greeting string
	Mu       sync.Mutex
	Chan     chan string
}
