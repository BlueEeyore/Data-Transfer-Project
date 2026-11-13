package send

import (
	"bufio"
	"log"
	"strings"
	"io"
)

func ReadNetInfo(infoStr string) (string, string) {
	reader := bufio.NewReader(strings.NewReader(infoStr))

	network, err := reader.ReadString('-')
	if err != nil {
		log.Fatalln("failed to read network", err)
	}
	network = strings.TrimSuffix(network, "-")

	ip, err := reader.ReadString('-')
	if err != nil {
		log.Fatalln("failed to read ip", err)
	}
	ip = strings.TrimSuffix(ip, "-")

	portBytes, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalln("failed to read port", err)
	}
	port := string(portBytes)
	
	return network, ip+":"+port
}
