package main

import "fmt"

func main() {
	var variabelA int = 4
	var variabelB *int = &variabelA

	fmt.Println("varA value: ", variabelA)
	fmt.Println("varA address: ", &variabelA)

	fmt.Println("varB value: ", *variabelB)
	fmt.Println("varB address: ", variabelB)
}
