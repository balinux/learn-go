// Closure jenis IIFE ini eksekusinya adalah langsung saat deklarasi.
// Teknik ini biasa diterapkan untuk membungkus proses yang hanya dilakukan sekali.
// IIFE bisa memiliki nilai balik atau bisa juga tidak.
package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4, 3, 2, 1, 2, 2, 2, 3, 4, 4, 5, 5}

	newNumbers := func(min int) []int {
		var r []int
		for _, v := range numbers {
			if v < min {
				continue
			}
			r = append(r, v)
		}
		return r
	}(3)

	fmt.Println("numbers: ", numbers)
	fmt.Println("new numbers: ", newNumbers)
}
