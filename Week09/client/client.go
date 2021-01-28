package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)
	log.Println("Client Start...")
	conn1, err := net.Dial("tcp", "127.0.0.1:8888")
	conn2, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Printf("Error dialing: %v\n", err.Error())
		return
	}
	group.Go(func() error {
		return receive(conn1)
	})

	group.Go(func() error {
		return sendMsg(conn1)
	})
	group.Go(func() error {
		return receive(conn2)
	})
	group.Go(func() error {
		return sendMsg(conn2)
	})

	select {
	case <-ctx.Done():
		conn1.Close()
		log.Println("conn1 closed!!!")
		conn2.Close()
		log.Println("conn2 closed!!!")
	}

	if err := group.Wait(); err != nil {
		cancel()
		log.Printf("error: %+v \n", err)
	}

}

func receive(conn net.Conn) error {
	rd := bufio.NewReader(conn)
	for {
		line, _, err := rd.ReadLine()
		fmt.Println(line)
		if err != nil {
			log.Printf("read error: %v\n", err)
			return err
		}
	}
}

func sendMsg(conn net.Conn) error {
	inputReader := bufio.NewReader(os.Stdin)
	outputWriter := bufio.NewWriter(conn)
	for {
		fmt.Println(conn.LocalAddr(), " Please write your send massage:")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return errors.New("quit")
		}
		_, _ = outputWriter.Write([]byte(trimmedInput + "\n"))
		_ = outputWriter.Flush()
	}
}
