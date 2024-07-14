package main

import "fmt"

func main() {
	// defer fmt.Println("hallo")
	// fmt.Println("ini duluan")

	// defer in function
	oderFood("pizza")
	oderFood("burger")

	// kombinasi defer dengan IIFE
	number := 3
	if number == 3 {
		fmt.Println("hallo 1")
		defer fmt.Println("hello 3")
	}

	fmt.Println("hello 2")

	// agar halo 3 bisa di eksekusi duluam maka harus di bungkus IIFE
	if number == 3 {
		fmt.Println("hallo 4")
		func() {
			defer fmt.Println("halo 6")
		}()
	}
	fmt.Println("halo 5")
}

func oderFood(menu string) {
	defer fmt.Println("terima kasih sudah menunggu")
	if menu == "burger" {
		fmt.Println("kamu order burger")
		return
	}
	fmt.Println("pesanan kamu", menu)
}
