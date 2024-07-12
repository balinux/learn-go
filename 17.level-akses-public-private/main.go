package main

import (
	initfunction "belajar-golang-level-akses/initFunction"
	"belajar-golang-level-akses/library"
	. "belajar-golang-level-akses/prefixgo"
	"belajar-golang-level-akses/student"
	"fmt"

	// alias
	f "fmt"
)

func main() {
	library.Sayhello("rio")

	// penerapan export struct
	s1 := student.Student{"rio", 22}
	f.Println("stuct exported")
	f.Println("hallo bro", s1.Name)
	f.Println("age", s1.Name)

	// perepan import prefix titil
	prefixdot := PrefixExport{Message: "wow"}
	f.Println("prefix import titik")
	f.Println("hallo prefixdot:", prefixdot.Message)

	// memanggil fungsi init()
	fmt.Println("Name from init:", initfunction.Student.Name)
	fmt.Println("Age from init:", initfunction.Student.Age)
}
