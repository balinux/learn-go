package main

import (
	"belajar-golang-level-akses/library"
	"belajar-golang-level-akses/student"
	"fmt"
)

func main() {
	library.Sayhello("rio")

	// penerapan export struct
	s1 := student.Student{"rio", 22}
	fmt.Println("stuct exported")
	fmt.Println("hallo bro", s1.Name)
	fmt.Println("age", s1.Age)
}
