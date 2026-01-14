package helpers

import "fmt"

// BroadCast a Message to all clients
func (s *Server) BroadcastMessages() {
	for msg := range s.Broadcast {
		// create copy of clients to not hold the mutex so long
		s.mu.Lock()
		clientsCopy := append([]*Client(nil), s.clients...)
		s.mu.Unlock()

		for _, c := range clientsCopy {
			if c == msg.client {
				continue
			}
			select {
			case c.send <- struct {
				message    string
				history    bool
				showPrompt bool
			}{message: msg.message, showPrompt: true}:
			default:
				fmt.Println("Kicking lagging client:", c.c_name)
				c.c_conn.Close()
			}
		}
	}
}
