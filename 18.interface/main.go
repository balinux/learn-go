package main

import (
	"fmt"
	e "interface/embeddedinterface"
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

	// mengakses method yang tidak terdefinisikan di interface
	// contoh yaitu method JariJari
	bangunDatarLingkaran := bangunDatar.(h.Lingkaran)
	fmt.Println("jari-jari: ", bangunDatarLingkaran.JariJari())

	// embedded interface
	var bangunRuang e.HitungEmbedded = &e.Kubus{Sisi: 30.0}
	fmt.Println("=== embedded interface kubud")
	fmt.Println("sisi:", bangunRuang)
	fmt.Println("keliling:", bangunRuang.Keliling())
	fmt.Println("luas:", bangunRuang.Luas())
	fmt.Println("volume:", bangunRuang.Volume())
}
