package main

import (
	"fmt"
	"github.com/BlueEeyore/Data-Transfer-Project.git/receive"
)

func main() {
	msg, err := receive.Receive("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
}
