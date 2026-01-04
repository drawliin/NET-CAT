package helpers

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn, ch chan string) {
	defer conn.Close()
	_, err := conn.Write([]byte(greeting))
	if err != nil {
		fmt.Println("error")
		return
	}
	//creating buffer
	buff := make([]byte, 1024)
	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(err)
		} else {
			//send message over the channel
			ch <- string(buff[:n])
		}
	}
}
