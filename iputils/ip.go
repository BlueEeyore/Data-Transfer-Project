package iputils

import (
	"bufio"
	"log"
	"strings"
	"io"
	"net"
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


func GetOutboundIp() (net.IP) {
	// send a dummy udp connection. Doesn't matter where it actually goes
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatalln("failed to send dummy udp", err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
