package main

import (
	"fmt"
	"os"
)

func main() {
	argsRaw := os.Args
	fmt.Println("-> %#v\n", argsRaw)

	args := argsRaw[1:]
	fmt.Println("->%#v\n", args)
}
