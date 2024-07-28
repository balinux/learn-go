package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// cara mendapakan working directory saat ini
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentDir)

	// membukan dan membaca file di folder saat ini
	files, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	// menghitung jumlah file berdasarkan extensi

	fileExtentions := make(map[string]int)
	fmt.Println("fileExtentions sebelum di isi", fileExtentions)
	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			fmt.Println("ekstensi", fileExtentions[ext])
			fileExtentions[ext]++
		}
	}
	fmt.Println("fileExtentions setelah di isi", fileExtentions)

	// menampilakn jumlah file
	fmt.Printf("jumlah file di dalam folder %s \n", currentDir)
	for ext, count := range fileExtentions {
		fmt.Printf("%s:%d\n", ext, count)
	}
}
