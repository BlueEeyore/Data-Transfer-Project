package send

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"log"
)

func Send(network, address, message string) (error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Printf("failed to connect to network %v and address %v", network, address)
		return err
	}
	defer conn.Close()

	fmt.Fprintf(conn, message)

	return nil
}

func SendInput(network, address string) {
	// connect to this address
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// scan stdin line by line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inp := scanner.Text()
		fmt.Fprintln(conn, inp)
	}

	// send "/over" to signal end of message
	fmt.Fprintln(conn, "/over")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
