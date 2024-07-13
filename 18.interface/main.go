package main

import (
	"fmt"
	h "interface/hitung"
)

func main() {
	var bangunDatar h.Hitung

	bangunDatar = h.Persegi{Sisi: 30.0}
	fmt.Println("==== persegi")
	fmt.Println("luas:", bangunDatar.Luas())
	fmt.Println("keliling:", bangunDatar.Keliling())

	bangunDatar = h.Lingkaran{Diameter: 22.9}
	fmt.Println("==== Lingkaran")
	fmt.Println("luas:", bangunDatar.Luas())
	fmt.Println("keliling", bangunDatar.Luas())
}
