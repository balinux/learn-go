package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Print("tuliskan sebuah kata:")
	fmt.Scanln(&input)

	var number int
	var err error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println("number is", number)
	} else {
		fmt.Println(number, "is not a number")
		fmt.Println(err.Error())
	}

	// membuat custom error
	var name string
	fmt.Print("masukkan nama:")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo:", name)
	} else {
		// contoh penggunaan panic
		// panic(err.Error()) // response : /Users/yhotie/Documents/code/go/learn-go/28.error-panic-recover.go:34 +0x38c
		fmt.Println("end")
	}
}

// membuat fungsi untuk handle custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cant not be empty")
	}
	return true, nil
}
