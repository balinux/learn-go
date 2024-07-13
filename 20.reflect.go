package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Mendapatkan tipe data dan nilai menggunakan reflect
	data := 12
	reflectData := reflect.ValueOf(data)

	fmt.Println("value data:", reflectData)
	fmt.Println("type variabel:", reflectData.Type())

	// Mengakses nilai dalam bentuk Interface{}
	fmt.Println("nilai variabel:", reflectData.Interface())

	// Mengambil nilainya
	nilaiReflectData := reflectData.Interface().(int)
	fmt.Println(nilaiReflectData)

	x1 := &Xxx{Name: "rio", Age: 33}
	x1.getPropertyValue()
}

// xxx struct untuk mendemonstrasikan reflect
type Xxx struct {
	Name string
	Age  int
}

// Akses informasi properti variabel objek
func (x *Xxx) getPropertyValue() {
	reflectValue := reflect.ValueOf(x)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama:  ", reflectType.Field(i).Name)
		fmt.Println("tipe data:  ", reflectType.Field(i).Type)
		fmt.Println("nilai:  ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
