package main

import "fmt"

func main() {
	dataSlice := []string{"rio", "juniyantara", "putra"}

	fmt.Println(dataSlice[1], "\n")

	// Hubungan Slice Dengan Array & Operasi Slice
	buah := []string{"aaa", "bbb", "ccc", "ddd"}
	buahBaru := buah[0:2]
	fmt.Println(buahBaru)

	// slice merupakan tipda data reference
	buahA := []string{"ddd", "fff", "ccc"}
	buahAa := buahA[0:2]
	buahAaA := buahAa[1:2]

	fmt.Println(buahA)
	fmt.Println(buahAa)
	fmt.Println(buahAaA)

	// fungsi lend dan cap
	// len untuk menghitung kapasitas maksimum slice
	// len untuk menghitung jumlah element

	fruits := []string{"zzz", "xxx", "ccc"}
	fmt.Println("len: ", len(fruits))
	fmt.Println("cap: ", cap(fruits))

	fruits1 := fruits[0:2]
	fruits2 := fruits[0:2:2]

	fmt.Println("fruits1: ", fruits1)
	fmt.Println("len1: ", len(fruits1))
	fmt.Println("cap1: ", cap(fruits1))

	fmt.Println(fruits2)
	fmt.Println("len2: ", len(fruits2))
	fmt.Println("cap2: ", cap(fruits2))
}
