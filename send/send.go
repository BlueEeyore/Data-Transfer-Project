package send

import (
	"fmt"
	"net"
)

func Send(network, address, message string) (error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Fprintf(conn, message)

	return nil
}
