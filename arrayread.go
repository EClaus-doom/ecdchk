package main

import (
	"io"
	"os"
)

// function example from stackoverflow thread below...
// https://stackoverflow.com/questions/70115047/convert-binary-file-to-array-of-bytes
// the function is written by both https://stackoverflow.com/users/17014998/omermichleviz
// and https://stackoverflow.com/users/1205448/dylan-reimerink
// seems valid as of Go 1.23.3. shortened the filepath name and function name to be shorter,
// also io.ReadAll is supposed to return something so I have it return "file"
// as file is the buffer, this is also technically a slice not an array
func readIntoByteSlice(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return io.ReadAll(file)
}
