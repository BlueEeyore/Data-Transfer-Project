package file_copy

import (
	"io"
	"os"
)

func Copy(in string, out string) (int64, error) {
	src, err := os.Open(in)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.Create(out)
	if err != nil {
		return 0, err
	}
	// don't defer close, as error may come from closing file

	nbytes, err := io.Copy(dst, src)
	if err != nil {
		return nbytes, err
	}
	
	err = dst.Close()
	if err != nil {
		return nbytes, err
	}

	return nbytes, nil
}
