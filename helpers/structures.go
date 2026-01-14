package helpers

import (
	"net"
	"sync"
)

type Message struct {
	client  *Client
	message string
}

type Server struct {
	clients   []*Client
	mu        sync.Mutex
	Broadcast chan Message
	history   []string
	Greeting  []byte
}

type Client struct {
	c_conn net.Conn
	c_name string
	send   chan struct {
		message    string
		history    bool
		showPrompt bool
	}
}
