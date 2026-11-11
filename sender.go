package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/send"
)

func main() {
	err := send.Send("tcp", "localhost:8080", "Hello, World!")
	if err != nil {
		panic(err)
	}
}
