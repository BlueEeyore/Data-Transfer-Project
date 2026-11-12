package send

import (
	"fmt"
	"net"
	"bufio"
	"os"
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

func SendInput(network, address string) (error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		inp := scanner.Text()
		fmt.Fprintln(conn, inp)
	}
	fmt.Fprintf(conn, "/over")

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
