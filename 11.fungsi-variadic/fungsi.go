package main

import "fmt"

func main() {
	average := calculate(2, 3, 42, 6, 7, 87, 65, 20)
	message := fmt.Sprintf("nilai rata-rata = %.2f\n", average)
	fmt.Println(message)

	// menggunakan data slice
	dataVariadic := []int{1, 2, 3, 4, 23, 3, 2, 2, 44, 55, 33, 55}
	averageSlice := variadicWithSlice(dataVariadic...)
	messageSlice := fmt.Sprintf("nilai rata-rata data slice: %.2f\n", averageSlice)
	fmt.Println(messageSlice)
}

// return rata-rata bilangan dari banyak bilangan
func calculate(numbers ...int) float64 {
	total := 0
	for _, value := range numbers {
		total += value
	}

	average := float64(total) / float64(len(numbers))

	return average
}

// membuatf fungsi untuk menampung data slice variadic
func variadicWithSlice(numbers ...int) float64 {
	totals := 0
	for _, value := range numbers {
		totals += value
	}

	average := float64(totals) / float64(len(numbers))

	return average
}
