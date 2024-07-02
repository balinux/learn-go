package main

import "fmt"

func main() {
	var positiveNumber uint8 = 64
	negativNumber := -123123

	fmt.Printf("bilangan positif %d \n", positiveNumber)
	fmt.Printf("bilangan negatif %d \n", negativNumber)

	// tipe data numerik decimal
	decimalNumber := 2.78
	fmt.Printf("bilangan desimal %f\n", decimalNumber)
	fmt.Printf("bilangan desimal %.3f\n", decimalNumber)

	// type data boolean
	var isTrue bool = true
	fmt.Printf("is data true: %t\n", isTrue)

	// type string
	messageMultiLine := `Nama saya "rio".
  salam kenal.
    saya mau belajar "golang"
    `
	fmt.Println(messageMultiLine)
}
