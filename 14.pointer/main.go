package main

import "fmt"

func main() {
	age := 30
	fmt.Println("usia awal", age)

	change(&age, 33)
	fmt.Println("usia setelah ganti", age)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("umur diganti dengan: ", *old)
}
