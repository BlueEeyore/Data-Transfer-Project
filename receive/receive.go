package receive

import (
	"net"
	"io"
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

