package initfunction

import "fmt"

var Student = struct {
	Name string
	Age  int
}{}

func init() {
	Student.Name = "rio"
	Student.Age = 33

	fmt.Println("initial imported component")
}
