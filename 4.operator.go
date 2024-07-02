package main

import "fmt"

func main() {
	// Operator Aritmatika
	value := (((2+6)%3)*4 - 2) / 3

	// Operator Perbandingan
	isEqual := (value == 2)

	fmt.Printf("nilai %d (%t)\n", value, isEqual)

	// Operator Logika
	left := false
	right := true

	leftAndRight := left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)

	leftOrRight := left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	leftReverse := !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
