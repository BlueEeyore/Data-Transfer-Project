package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/send"
)

func main() {
	send.SendInput("tcp", "localhost:8080")
}
