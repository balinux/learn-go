package main

import (
	"fmt"
	"io"
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

	// mendapatkan nama file yang desang di jalankan
	executable, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	executableName := filepath.Base(executable)

	// membukan dan membaca file di folder saat ini
	files, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatal(err)
	}

	// memindahkan file ke folder berdasarkan ekstensi
	for _, file := range files {
		if !file.IsDir() && file.Name() != executableName {
			ext := filepath.Ext(file.Name())
			if ext == "" {
				fmt.Println("file has no extention")
				// jika tidak memiliki ekstensi maka berikan nama folder no_extension
				ext = "no_extension"
			} else {
				fmt.Println("ekstensi file", ext[1:])
				ext = ext[1:]
			}

			// membaut folder berdasarkan ekstensi jika belum ada
			extDir := filepath.Join(currentDir, ext)
			if _, err := os.Stat(extDir); os.IsNotExist(err) {
				err := os.Mkdir(extDir, 0755)
				if err != nil {
					log.Fatal(err)
				}
			}

			// memindahkan file ke folder yang sesuai berdasarkan path
			oldPath := filepath.Join(currentDir, file.Name())
			newPath := filepath.Join(extDir, file.Name())
			err := moveFile(oldPath, newPath)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// memindahkan file dari oldpath ke new path
func moveFile(oldPath, newPath string) error {
	inputFile, err := os.Open(oldPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	outputFile, err := os.Create(newPath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// bagian untuk mengcopy file
	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	// hapus file
	err = os.Remove(oldPath)
	if err != nil {
		return err
	}
	return nil
}
