package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "rio"
	names[1] = "juniyantara"
	names[2] = "putra"

	fmt.Println(names[0], names[1], names[2])

	// Inisialisasi Nilai Awal Array
	frts := [3]string{"manggo", "apple", "orange"}
	fmt.Println("Jumlah data", len(frts))
	fmt.Println("isi semua element", frts)

	//  Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	frts2 := [...]string{"aaa", "sss", "dddd", "gggg"}
	fmt.Println("jumlah data", len(frts2))
	fmt.Println("isi semua element", frts2)

	// array multidimensi
	number1 := [2][3]int{{3, 4, 5}, {4, 4, 5}}
	number2 := [2][3]int{{4, 5, 6}, {7, 8, 9}}
	fmt.Println("number1: ", number1)
	fmt.Println("number2: ", number2)

	// Perulangan Elemen Array Menggunakan Keyword for
	number3 := [4]string{"asd", "ddd", "ccc", "sss"}

	for i := 0; i < len(number3); i++ {
		fmt.Printf("element %d:%s\n", i, number3[i])
	}

	//  Perulangan Elemen Array Menggunakan Keyword for - range
	number4 := [4]string{"asd", "ddd", "ccc", "sss"}
	for i, data := range number4 {
		fmt.Printf("element for range %d: %s\n", i, data)
	}

	// Penggunaan Variabel Underscore _ Dalam for - range
	number5 := [4]string{"asd", "ddd", "ccc", "sss"}
	for _, data := range number5 {
		fmt.Printf("element for range Underscore  %s\n", data)
	}

	// Alokasi Elemen Array Menggunakan Keyword make
	namaBuah := make([]string, 2)
	namaBuah[0] = "apple"
	namaBuah[1] = "jeruk"
	fmt.Println(namaBuah)
}
