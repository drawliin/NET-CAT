package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 1024)
	fmt.Println("server reading...")
	n, err := conn.Read(buff) //blocking
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("server successfuly read:")
	fmt.Println(string(buff[:n]))
	fmt.Println("server writing...")
	fmt.Fprint(conn, "echo - "+string(buff[:n]))
}

func server(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	listener, err := net.Listen("tcp", "localhost:8989")
	if err != nil {
		log.Println("error listening")
		return
	}
	fmt.Println("server is listening")
	ch <- struct{}{}
	defer listener.Close()
	fmt.Println("accept is run next")
	for {
		//accept incoming connections
		conn, err := listener.Accept() //blocking
		fmt.Println("client conn retrieved")
		if err != nil {
			fmt.Println(err)
			continue
		}
		//handle connection
		go handleConnection(conn)
	}
}

func client(wg *sync.WaitGroup, ch chan struct{}) {
	defer wg.Done()
	<-ch
	conn, err := net.Dial("tcp", "localhost:8989")
	if err != nil {
		fmt.Println("client: error connecting")
		return
	}
	fmt.Println("client connected")
	defer conn.Close()
	//send message
	fmt.Fprint(conn, "hello there")
	fmt.Println("message sent to server")
	buff := make([]byte, 1024)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Println("client: error reading")
	}
	fmt.Println("message received from server:", string(buff[:n]))

}
func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{})
	//server
	wg.Add(1)
	go server(&wg, ch)
	//client
	wg.Add(1)
	go client(&wg, ch)
	//
	wg.Wait()
	fmt.Println("all workers finished")
}
