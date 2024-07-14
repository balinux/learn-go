package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	text := "secret"
	shaSecret := sha1.New()
	shaSecret.Write([]byte(text))
	encrypted := shaSecret.Sum(nil)
	encryptedString := fmt.Sprintf("%x", encrypted)
	fmt.Println(encryptedString)

}
