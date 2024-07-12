package main

import "fmt"

func main() {
	s1 := student2{"rio Juniyantara", 22}
	fmt.Println("s1 before: ", s1.name)

	s1.changeName1("baba yaga")
	fmt.Println("s1 after changeName1", s1.name)

	s1.changeName2("puta blyat")
	fmt.Println("s1 after changeName2", s1.name)
}

func (s student2) changeName1(name string) {
	fmt.Println(" ---> in changeName1, name chaged to:", name)
	s.name = name
}

func (s *student2) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed tp:", name)
	s.name = name
}

type student2 struct {
	name string
	age  int
}
