package main

import "fmt"

func main() {
	getMinMax := func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e

			}
		}
		return min, max
	}

	numbers := []int{2, 5, 3, 4, 4, 3, 5, 2, 5}
	min, max := getMinMax(numbers)

	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max)
}
