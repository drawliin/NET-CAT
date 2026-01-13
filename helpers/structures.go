package helpers

import (
	"net"
	"sync"
)

type Server struct {
	clients   []*Client
	mu        sync.Mutex
	Broadcast chan string
	history   []string
	Greeting  []byte
}

type Client struct {
	c_conn net.Conn
	c_name string
	send   chan string
}
