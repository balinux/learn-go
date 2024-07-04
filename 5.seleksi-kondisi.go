package main

import "fmt"

func main() {
	var nilai int8 = 8

	if nilai == 10 {
		fmt.Println("anda lulus dengan sempurna")
	} else if nilai > 7 {
		fmt.Print("anda lulus\n")
	} else if nilai < 7 {
		fmt.Println("anda tidak lulus")
	} else {
		fmt.Printf(" anda tidak lulus dengan nilai %d\n", nilai)
	}
}
