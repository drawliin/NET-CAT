package helpers

import "bufio"

func (s *ChatServer) broadCastMessages() {
	for {
		message := <-s.MsgChan
		//loop over connections and broadcast message
		s.Mu.Lock()
		for _, conn := range s.Clients {
			writer := bufio.NewWriter(conn)
			writer.WriteString(message)
			writer.Flush()
		}
		s.Mu.Unlock()
	}
}
