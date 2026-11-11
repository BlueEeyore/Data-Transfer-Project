package file_copy

import (
	"io"
	"os"
)

func Copy(in string, out string) (bytes int64, err error) {
	src, err := os.Open(in)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.Create(out)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	bytes, err = io.Copy(dst, src)
	if err != nil {
		return 0, err
	}

	return bytes, nil
}
