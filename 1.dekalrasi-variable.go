package main

import "fmt"

func main() {
	var firstName string = "rio"
	fmt.Printf("Selamat datang %s! \n", firstName)
	fmt.Println("Selamat datang", firstName+"!")

	// manifest typing
	var fName string = "rio"

	// type interface
	mName := "juniyantara"
	lName := "putra"

	fmt.Printf("selamat datang %s %s %s \n", fName, mName, lName)

	// dekalrasi multiple variable
	// mainfest typing
	var a, b, c string
	a, b, c = "rio", "juniyantara", "putra"

	// type interface
	one, two, tree := "one", 2, 3
	four, five, six := "4", 5, "6"
	fmt.Printf("hallo %s %s %s \n", a, b, c)
	fmt.Printf("hallo %s %d %d \n", one, two, tree)
	fmt.Printf("hallo %s %d %s \n", four, five, six)

	// variabel under score (_)
	_ = "under score var"
	// fmt.Println("tes print under score", _)

	// Dekalari variabel menggunakan keyword new(type) pointer
	typeNew := new(string)
	*typeNew = "nilai pointer"
	fmt.Println(typeNew)
	fmt.Println(*typeNew)
}
