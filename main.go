package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/receive"
	"github.com/BlueEeyore/Data-Transfer-Project.git/send"
	"github.com/BlueEeyore/Data-Transfer-Project.git/iputils"
	"flag"
	"fmt"
)

func main() {
	sending := flag.Bool("send", false, "sending a file")
	fn := flag.String("fn", "", "the filename of the image to be sent")
	nw := flag.String("nw", "", "the network info (protocol and address")
	protocol := flag.String("p", "tcp", "the network protocol")
	port := flag.String("port", "8080", "the port that will be used")

	flag.Parse()

	if *sending {
		if *fn == "" {
			fmt.Println("please specify filename with -fn flag")
		} else {
			ip := iputils.GetOutboundIp().String()

			fmt.Printf("%v-%v-%v\n", *protocol, ip, *port)

			send.SendFile(*protocol, ip+":"+*port, *fn)
		}
	} else {
		if *nw == "" {
			fmt.Println("please provide network information")
		} else {
			prot, address := iputils.ReadNetInfo(*nw)

			receive.ReceiveFile(prot, address)
		}
	}
}
