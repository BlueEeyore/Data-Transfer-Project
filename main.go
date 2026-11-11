package main

import (
	"github.com/BlueEeyore/Data-Transfer-Project.git/file_copy"
)

func main() {
	file_copy.Copy("hello.txt", "backup.txt")
}
