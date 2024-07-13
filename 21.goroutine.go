package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	go print(15, "hallo")
	print(15, "tes goroutine")

	var input string
	fmt.Scanln(&input)
}
