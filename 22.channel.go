package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// membuat channel
	message := make(chan string)

	sayHello := func(who string) {
		data := fmt.Sprintf("hello %s", who)
		// pengiriman data ke channel message
		message <- data
	}

	// penerapan beberapa goroutine
	// eksekusi goroutine adalah asynchronous, sedangkan serah-terima data antar channel adalah synchronous.
	go sayHello("Rio")
	go sayHello("Rian")
	go sayHello("Linda")

	message1 := <-message
	fmt.Println(message1)

	message2 := <-message
	fmt.Println(message2)

	message3 := <-message
	fmt.Println(message3)

	// channel sebagai tipe data parameter
	messages := make(chan string)

	for _, each := range []string{"sss", "aaa", "ggg"} {
		go func(who string) {
			data := fmt.Sprintf("channel as parameter: %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
