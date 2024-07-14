package main

import (
	"flag"
	"fmt"
)

func main() {
	// data1 := flag.String("name", "anonymous", "masukkan anama anda")
	// fmt.Println(*data1)

	var data2 string
	flag.StringVar(&data2, "name", "anonym", "masukkan anma 2")
	fmt.Println(data2)
}
