package receive

import (
	"net"
	"io"
	"bufio"
	"log"
	"fmt"
	"os"
	"github.com/BlueEeyore/Data-Transfer-Project.git/file"
	"encoding/json"
	"strings"
	"errors"
	"strconv"
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


func ReceiveFile(network, address string) {
	// connect to this address
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// decode received json data into FileMessage format
	var fm file.FileMessage
	err = json.NewDecoder(conn).Decode(&fm)
	if err != nil {
		log.Fatalln("failed to json decode", err)
	}

	dstName := uniqueDstName(fm.Name)

	// write bytes into new file with same name as received
	// 0644 sets standard file permissions (owner: read/write, others: read-only)
	err = os.WriteFile(dstName, fm.File, 0644)
	if err != nil {
		log.Fatal(err)
	}

	conn.Close()
}

func uniqueDstName(name string) (string) {
	// produce a filename that does not already exist in current directory
	lastPeriod := strings.LastIndex(name, ".")
	dstPrefix := name[:lastPeriod]
	dstName := dstPrefix + name[lastPeriod:]
	i := 0
	for {
		fmt.Println(dstName)
		_, err := os.Stat(dstName)
		if errors.Is(err, os.ErrNotExist) {
			break
		} else if err != nil {
			log.Fatalln("failed to check for existing file name", err)
		}
		dstName = dstPrefix + strconv.Itoa(i) + name[lastPeriod:]
		i++
	}
	return dstName
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
