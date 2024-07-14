package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	timeout := 5
	ch := make(chan bool)

	//jalankan go routine
	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string
	fmt.Println("berapa 2+3?")
	fmt.Scanln(&input)
	if input == "5" {
		fmt.Println("the answer is right!")
	} else {
		fmt.Println("the answer is wrong!")
	}
}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("time out", timeout, "Second")
	os.Exit(0)
}
