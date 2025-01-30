package utilities

import (
	"errors"
	"os"
)

func ReadInput() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New("usage: go run main.go <input_file_name>")
	}
	return os.Args[1], nil
}
