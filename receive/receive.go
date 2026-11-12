package receive

import (
	"net"
	"io"
	"bufio"
)

func Receive(network, address string) (string, error) {
	ln, err := net.Listen(network, address)
	if err != nil {
		return "", err
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		return "", err
	}

	msg, err := io.ReadAll(conn)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}

func HandleReceive(network, address string, ch chan string) {
	ln, err := net.Listen(network, address)
	if err != nil {
		panic(err)	// handle errors
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)	//handle errors
		}

		go handleConn(conn, ch)
	}
	close(ch)
}


func handleConn(conn net.Conn, ch chan string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "/over" {
			break
		}

		ch<-line
	}
}
