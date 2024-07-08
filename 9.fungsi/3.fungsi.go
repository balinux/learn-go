package main

import "fmt"

func main() {
	dividerNumber(20, 2)
	dividerNumber(12, 0)
	dividerNumber(32, 5)
}

// penggunakan keyword 'return' untuk menghentikan proses dalam fungsi
func dividerNumber(a, b int) {
	if b == 0 {
		fmt.Printf("data %d tidak bisa di bagi dengan %d\n", a, b)
		return
	}

	hasil := a / b
	fmt.Printf("hasil bagi %d dengan %d adalah %d\n", a, b, hasil)
}
