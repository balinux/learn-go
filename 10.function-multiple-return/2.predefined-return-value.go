package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	// menampung data kemalian multiple return
	area, circumference := calculateData(diameter)

	fmt.Printf("luas lingkaran: %.2f\n", area)
	fmt.Printf("keliling lingkaran: %.2f\n", circumference)
}

// sebuah fungsi untuk menghitung luas, dan keliling lingkaran
// predefiner bisa mendefinisikan nilai balik lebih awal sepeti conth di bawah ini
func calculateData(d float64) (area float64, circumference float64) {
	// mengihitung luas lingkaran
	area = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling lingkaran
	circumference = math.Pi * d

	return
}
