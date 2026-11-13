package send

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"log"
	"github.com/BlueEeyore/Data-Transfer-Project.git/file"
	"encoding/json"
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

func SendFile(network, address, inpFile string) {
	// initialise a listener to address
	ln, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	// read file bytes of inpFile
	srcbytes, err := os.ReadFile(inpFile)
	if err != nil {
		log.Fatal(err)
	}

	// initialise FileMessage type
	fm := file.NewFile(inpFile, srcbytes)

	// convert to json
	jbytes, err := json.Marshal(fm)
	if err != nil {
		log.Fatal(err)
	}

	// wait until someone dials and accept it
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	// write json bytes to connection
	_, err = conn.Write(jbytes)
	if err != nil {
		log.Fatalln("failed to write bytes to connection", err)
	}
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
