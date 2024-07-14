package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(randomizer.Int()%10+1) * time.Second)
	}
}

func retreiveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Println("rereive data:", data)
		case <-time.After(time.Second * 5):
			fmt.Println("timeout, durasi melebihi 5 detik")
			break loop
		}
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	message := make(chan int)
	go sendData(message)
	retreiveData(message)
}
