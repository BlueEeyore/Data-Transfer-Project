package receive

import (
	"net"
	"log"
)

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
