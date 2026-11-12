package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/receive"
)

func main() {
	receive.Receive("tcp", ":8080")
}
