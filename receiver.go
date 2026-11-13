package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/receive"
	"fmt"
)

func main() {
	network := "tcp"
	ip := receive.GetOutboundIp().String()
	port := "8080"

	fmt.Println(network, ip, port)

	receive.ReceiveFile(network, ip+":"+port)
}
