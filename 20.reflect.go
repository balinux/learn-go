package main

import (
	"fmt"
	"reflect"
)

func main() {
	// mendari type data dan value menggunakan reflect
	data := 12
	reflectData := reflect.ValueOf(data)

	fmt.Println("value data:", reflectData)
	fmt.Println("type variabel:", reflectData.Type())

	// megnakses nilai dalam bentuk Interface{}
	fmt.Println("nilai variabel:", reflectData.Interface())
	// mengambil nilainya
	nilaireflectData := reflectData.Interface().(int)
	fmt.Println(nilaireflectData)

	x1 := &xxx{Name: "rio", Age: 33}
	x1.getPropertyValue()
}

type xxx struct {
	Name string
	Age  int
}

// peraksesan informasi properti variabel objek
func (x *xxx) getPropertyValue() {
	reflectValue := reflect.ValueOf(x)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	reflectType := reflectValue.Type()

	for i := 0; i <= reflectValue.NumField(); i++ {
		fmt.Println("nama:  ", reflectType.Field(i).Name)
		fmt.Println("tipe data:  ", reflectType.Field(i).Type)
		fmt.Println("nilai:  ", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}
