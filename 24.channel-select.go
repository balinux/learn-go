package main

import (
	"fmt"
	"runtime"
)

func getAvg(numbers []int, ch chan float64) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	max := numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}
	ch <- max
}

func main() {
	runtime.GOMAXPROCS(2)
	numbers := []int{3, 6, 5, 7, 4, 2, 4, 8, 9, 6, 7, 4}
	fmt.Println("numbers:", numbers)

	ch1 := make(chan float64)
	go getAvg(numbers, ch1)

	ch2 := make(chan int)
	go getMax(numbers, ch2)

	// tempat untuk menampung data
	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Println("nilai rata-rata:", avg)

		case max := <-ch2:
			fmt.Println("Max number:", max)
		}
	}
}
