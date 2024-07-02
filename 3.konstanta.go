package main

import "fmt"

func main() {
	const firstName string = "rio"
	const lastName = "juniyantara"
	fmt.Print("halo ", firstName, " ", lastName, "\n")

	// Deklarasi Multi Konstanta
	const (
		square        = "kotak"
		isToday  bool = true
		floatNum      = 2.2
	)

	// Ketika tipe data dan nilai tidak dituliskan dalam deklarasi konstanta,
	// maka tipe data dan nilai yang dipergunakan adalah sama seperti konstanta yang dideklarasikan diatasnya.
	const (
		a = "konstanta"
		b
	)
}
