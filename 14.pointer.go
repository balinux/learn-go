package main

import "fmt"

func main() {
	var variabelA int = 4
	var variabelB *int = &variabelA

	fmt.Println("varA value: ", variabelA)
	fmt.Println("varA address: ", &variabelA)

	fmt.Println("varB value: ", *variabelB)
	fmt.Println("varB address: ", variabelB)

	// efek perubahan nilai pointer
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("nilai numberA", numberA)
	fmt.Println("address numberA", &numberA)
	fmt.Println("nilai numberB", *numberB)
	fmt.Println("nilai numberB", numberB)

	fmt.Println("")
	numberA = 6

	fmt.Println("nilai numberA", numberA)
	fmt.Println("address numberA", &numberA)
	fmt.Println("nilai numberB", *numberB)
	fmt.Println("nilai numberB", numberB)

	// pointer sebagai parameter fungsi
	var numberParam int = 4
	fmt.Println("number before change: ", numberParam)

	change(&numberParam, 10)
	fmt.Println("number after change: ", numberParam)
}

func change(original *int, value int) {
	*original = value
}
