package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "anonymous", "type your name")
	age := flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Println("name:", *name)
	fmt.Println("age:", *age)
}
