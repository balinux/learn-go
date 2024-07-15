package main

import (
	"fmt"
	"os"
)

var path = "/Users/yhotie/Documents/code/go/learn-go/36.file/tes.txt"

func main() {
	createFile()
}

// handle error
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile() {
	// deteksi apakah file sudah ada
	_, err := os.Stat(path)

	// buat file baru
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("file berhasil di buat", path)
}
