package receive

import (
	"net"
	"io"
	"bufio"
	"log"
	"fmt"
)

func ReceivePing(network, address string) (string, error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		log.Printf("failed to listen to network %v and address %v", network, address)
		return "", err
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Println("failed to accept new client")
		return "", err
	}

	msg, err := io.ReadAll(conn)
	if err != nil {
		log.Println("failed to read from connection")
		return "", err
	}

	return string(msg), nil
}


func Receive(network, address string) {
	ch := make(chan string)
	go HandleReceive(network, address, ch)
	
	// print messages received on this address
	for msg := range ch {
		fmt.Println(msg)
	}
}

func HandleReceive(network, address string, ch chan string) {
	// make a listener for this address
	ln, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		// wait for new connection and accept
		conn, err := ln.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, ch)
	}
}


func handleConn(conn net.Conn, ch chan string) {
	// close connection after handled
	defer conn.Close()

	// read connection line by line and send to channel
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "/over" {
			break
		}

		ch<-line
	}
}
