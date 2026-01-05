package helpers

import (
	"net"
	"sync"
)

type ChatServer struct {
	Clients map[string]net.Conn
	Messages []string
	Greeting string
	Mu sync.Mutex
	MsgChan chan string
}