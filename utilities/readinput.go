package utilities

import (
	"errors"
	"os"
)

func ReadInput() (string, error) {
	args := os.Args
	if len(args) < 2 {

		return " ", errors.New("usage: go run main.go <input_file_name>")
	}

	return args[1], nil
}
