package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/send"
)

func main() {
	err := send.SendInput("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
}
