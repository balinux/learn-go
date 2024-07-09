package main

import "fmt"

func main() {
	max := 3
	numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	howMany, getNumbers := findMax(numbers, max)
	theNumbers := getNumbers()

	fmt.Println("numbers\t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3]
}

// Tentang fungsi findMax() sendiri, fungsi ini dibuat untuk mempermudah pencarian angka-angka yang nilainya di bawah atau sama dengan angka tertentu. Fungsi ini mengembalikan dua buah nilai balik:
//    Nilai balik pertama adalah jumlah angkanya.
//    Nilai balik kedua berupa closure yang mengembalikan angka-angka yang dicari.

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, v := range numbers {
		if v <= max {
			res = append(res, v)
		}
	}

	return len(res), func() []int {
		return res
	}
}
