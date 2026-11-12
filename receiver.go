package main

import (
	"fmt"
	"github.com/BlueEeyore/Data-Transfer-Project.git/receive"
)

func main() {
	ch := make(chan string)
	go receive.HandleReceive("tcp", ":8080", ch)
	for msg := range ch {
		fmt.Println(msg)
	}
}
