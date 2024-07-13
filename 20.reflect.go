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

	//
	//mengakses informasi method variabel object dengan reflect
	bakso := &Food{Name: "bakso", Taste: "good"}
	fmt.Println("nama:", bakso.Name)

	reflectBakso := reflect.ValueOf(bakso)
	methodBakso := reflectBakso.MethodByName("SetName")
	methodBakso.Call([]reflect.Value{
		reflect.ValueOf("sossis"),
	})

	fmt.Println("nama makanan baru:", bakso.Name)
}

type Food struct {
	Name  string
	Taste string
}

func (f *Food) SetName(name string) {
	f.Name = name
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
