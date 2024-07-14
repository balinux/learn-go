package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	time.Sleep(time.Second * 1)
	fmt.Println("ater 4 second")

	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	// fungsi time.Afterfunnc
	// pertama buat channel penalpung
	ch := make(chan bool)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	// fungsi time.after
	<-time.After(1 * time.Second)
	fmt.Println("expired time after")
}
