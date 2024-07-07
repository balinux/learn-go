package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// penggunaan keyword  for  edngan argument hanya kondisi
	b := 0
	for b < 5 {
		fmt.Println("for dengan argument kondisi:", b)
		b++
	}

	// penggunaan keyword for tarnpa argument
	a := 0

	for {
		fmt.Println("for tanpa argument angka", a)
		a++
		if a == 5 {
			break
		}
	}

	// penggunaan keyword for dan range
	xs := "123" // string
	for i, v := range xs {
		fmt.Println("index=", i, "value=", v)
	}

	ys := [5]int{20, 30, 40, 50, 60} // array
	for _, v := range ys {
		fmt.Println("value=", v)
	}

	// slice
	zs := ys[0:3]
	for _, v := range zs {
		fmt.Println("slice value=", v)
	}

	// map
	kvs := map[byte]int{'a': 0, 'b': 2, 'c': 33}
	for k, v := range kvs {
		fmt.Println("map key=", k, "value=", v)
	}
}
