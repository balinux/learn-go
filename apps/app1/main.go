package main

import (
	"fmt"
	"log"
	"os"
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

	// menghitung jumlah file (bukan folder)
	fileCount := 0
	for _, file := range files {
		if !file.IsDir() {
			fileCount++
		}
	}

	// menampilakn jumlah file
	fmt.Printf("jumlah file di dalam folder %s: %d\n", currentDir, fileCount)
}
