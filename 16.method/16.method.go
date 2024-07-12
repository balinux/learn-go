package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := student{"rio juniyantara putra", 22}
	s1.sayHello()

	name := s1.getNameAt(2)
	fmt.Println("hallo: ", name)
}

func (s student) sayHello() {
	fmt.Println("hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

type student struct {
	name string
	age  int
}
