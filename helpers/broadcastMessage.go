package helpers

import (
	"fmt"
	"net"
	"sync"
)

func BroadCastMessage(connections []net.Conn, mu *sync.Mutex, ch <-chan string) {
	for {
		message := <-ch
		mu.Lock()
		for _, conn := range connections {
			_, err := conn.Write([]byte(message))
			if err != nil {
				fmt.Println(err)
			}
		}
		mu.Unlock()
	}
}
