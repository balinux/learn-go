package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)
}

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func randomWithRange(min, max int) int {
	value := randomizer.Int()%(max-min+1) + min
	return value
}
