package library

import "fmt"

func Sayhello(name string) {
	fmt.Println("hello from library")
	introduce(name)
}

// function unexpoerted/private, hanya bisa di akses oleh internal saja
func introduce(name string) {
	fmt.Println("nama saya", name)
}
