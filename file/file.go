package file

type FileMessage struct {
	Name string
	File []byte
}

func NewFile(name string, f []byte) (FileMessage) {
	return FileMessage{Name: name, File: f}
}
