package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	// menampung data kemalian multiple return
	area, circumference := calculate(diameter)

	fmt.Printf("luas lingkaran: %.2f\n", area)
	fmt.Printf("keliling lingkaran: %.2f\n", circumference)
}

// sebuah fungsi untuk menghitung luas, dan keliling lingkaran
func calculate(d float64) (float64, float64) {
	// mengihitung luas lingkaran
	area := math.Pi * math.Pow(d/2, 2)

	// menghitung keliling lingkaran
	circumference := math.Pi * d

	return area, circumference
}
