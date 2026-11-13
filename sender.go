package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/send"
	"fmt"
)

func main() {
	network, address := send.ReadNetInfo()

	fmt.Printf("enter filename: ")
	var fn string
	fmt.Scanln(&fn)

	fmt.Println(network, address, fn)

	send.SendFile(network, address, fn)
}
