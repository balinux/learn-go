package main

import (
	"fmt"
	"os"
)

func main() {
	writeFile()
}

// handle error
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func writeFile() {
	path := "/Users/yhotie/Documents/code/go/learn-go/36.file/tes.txt"

	// buka file dengan akses read dan write
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// tulis data ke file
	_, err = file.WriteString("hallo\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("mari belajar golang\n")
	if isError(err) {
		return
	}

	// simpan perubahan
	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("file ebrhasil di isi")
}
