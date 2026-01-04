package main

import (
	"fmt"
	"net-cat/helpers"
	"os"
)

func main() {
	port := "8989"
	//validate args
	if len(os.Args) == 2 {
		port = fmt.Sprintf("%d", os.Args[1])
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	add := "localhost:" + port
	helpers.Listen(add)
}
