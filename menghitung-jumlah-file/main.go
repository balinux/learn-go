package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentDir)

	files, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)

	fileCount := 0
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println(file)
			fileCount++
		}
	}

	fmt.Printf("junlah file di dalam folder %s: %d\n", currentDir, fileCount)
}
