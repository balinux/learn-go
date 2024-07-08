package main

import (
	"fmt"
	"strings"
)

func main() {
	name := []string{"rio", "juniyantara"}
	printMessage("hello", name)
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
