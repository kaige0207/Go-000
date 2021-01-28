package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	log.Println("Server Start...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatalf("listen error: %+v\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %+v\n", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	msg := make(chan string, 1)
	defer conn.Close()
	defer close(msg)
	go readMsg(conn, msg)
	go receive(conn, msg)
	select {}
}

func readMsg(conn net.Conn, message <-chan string) {
	for msg := range message {
		fmt.Println(conn.RemoteAddr(), ":", msg)
		_, err := conn.Write([]byte("reviewed" + msg))
		if err != nil {
			log.Printf("write err %+v\n", err)
			return
		}
	}
}

func receive(conn net.Conn, message chan string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message <- scanner.Text()
	}
}
