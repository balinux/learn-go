package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"rio", "juniyantara", "putra"}

	dataMengandungHurufO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	dataPanjangHurufDiatas3 := filter(data, func(each string) bool {
		return len(each) > 3
	})

	fmt.Println("data asli \t\t:", data)

	fmt.Println("filter kata yang ada huruf o \t\t:", dataMengandungHurufO)

	fmt.Println("kata yang panjang diatas \"3\"\t:", dataPanjangHurufDiatas3)
}

// Skema closure bisa dijadikan alias dengan cara
type FilterCallback func(string) bool

// func filter(data []string, callback FilterCallback) []string {
// 	var result []string
// 	for _, each := range data {
// 		if filtered := callback(each); filtered {
// 			result = append(result, each)
// 		}
// 	}
// 	return result
// }

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
