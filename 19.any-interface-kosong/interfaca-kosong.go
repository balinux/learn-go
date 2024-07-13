package main

import (
	"fmt"
	"strings"
)

func main() {
	// interface kosong bisa menampung semua jenis data
	var secret interface{}

	// bisa di sis engan string
	secret = "rio"
	fmt.Println("secret: ", secret)

	secret = []string{"aaa", "bb", "ssdd"}

	// contoh penggunaan pada array
	data := map[string]interface{}{
		"name":  "rio",
		"age":   22,
		"hobby": []string{"run", "walk", "workout"},
	}

	// menggunakan any
	dataAny := map[string]interface{}{
		"aaa": "aaa",
		"bbb": 22,
	}

	fmt.Println(data)
	fmt.Println(dataAny)

	// casting variabel any
	var dataCasting interface{}
	dataCasting = 2
	// eprlu di casting ke tipe integer terlebih dahulu
	number := dataCasting.(int) * 10
	fmt.Println(dataCasting, " dikalikan engan 10: ", number)

	dataCasting = []string{"sss", "sdasd"}
	sambung := strings.Join(dataCasting.([]string), ", ")
	fmt.Println(sambung)

	// casting variabel interface kosong ke object pointer
	type secretstruct struct {
		name string
		age  int
	}

	var secretData interface{} = &secretstruct{name: "rio", age: 22}
	name := secretData.(*secretstruct).name
	fmt.Println("secret interface:", name)

	// kobinasi slice, map, interface kosong
	dataSlice := []map[string]interface{}{
		{"name": "aaa", "age": 33},
		{"name": "vaa", "age": 55},
		{"name": "vga", "age": 44},
	}

	for _, v := range dataSlice {
		fmt.Println("name:", v["name"])
	}
}
