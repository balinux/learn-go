package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	readFile()
}

func readFile() {
	path := "/Users/yhotie/Documents/code/go/learn-go/36.file/tes.txt"

	// buka file
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// baca file
	text := make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}

	fmt.Println("file berhasil di baaca")
	fmt.Println(string(text))
}

// handle error
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
