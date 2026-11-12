package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/receive"
	"github.com/BlueEeyore/Data-Transfer-Project.git/send"
	"time"
)

func main() {
	network, address := "tcp", "localhost:8081"
	go receive.Receive(network, address)
	time.Sleep(time.Second) // quick and dirty fix to listener not being ready in time
	send.SendInput(network, address)
}
