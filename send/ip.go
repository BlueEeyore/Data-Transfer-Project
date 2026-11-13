package send

import (
	"bufio"
	"os"
	"log"
	"fmt"
	"strings"
)

func ReadNetInfo() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter special string: ")

	network, err := reader.ReadString(' ')
	if err != nil {
		log.Fatalln("failed to read network from stdin", err)
	}
	network = strings.TrimSuffix(network, " ")

	ip, err := reader.ReadString(' ')
	if err != nil {
		log.Fatalln("failed to read ip from stdin", err)
	}
	ip = strings.TrimSuffix(ip, " ")

	portBytes, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("failed to read port from stdin", err)
	}
	port := string(portBytes)
	port = strings.TrimSuffix(port, "\n")
	
	return network, ip+":"+port
}
