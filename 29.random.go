package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomizer := rand.New(rand.NewSource(10))
	fmt.Println("random ke-1:", randomizer.Int())
	fmt.Println("random ke-2:", randomizer.Int())
	fmt.Println("random ke-3:", randomizer.Int())
}
