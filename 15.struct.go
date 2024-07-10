package main

import (
	"fmt"
)

func main() {
	type student struct {
		name  string
		grade int
	}

	var s1 student
	s1.name = "rio"
	s1.grade = 2

	fmt.Println("name: ", s1.name, " grade: ", s1.grade)

	// Inisialisasi object struct
	// ada 3 cara Inisialisasi struct
	var Inisialisasi1 student
	Inisialisasi1.name = "aaa"
	Inisialisasi1.grade = 200

	// cara kedua
	Inisialisasi2 := student{"bbb", 333}

	// cara ketiga
	Inisialisasi3 := student{name: "rio"}
	fmt.Println("inisialisasai struct")
	fmt.Println("inisialisasai cara 1: ", Inisialisasi1.name)
	fmt.Println("inisialisasai cara 2: ", Inisialisasi2.name)
	fmt.Println("inisialisasai cara 3: ", Inisialisasi3.name)

	// Varianel object pointer
	structPointer := student{name: "rio", grade: 1000}

	var struct2pointer *student = &structPointer

	fmt.Println("structPointer: ", structPointer.name)
	fmt.Println("structPointer2: ", struct2pointer.name)

	struct2pointer.name = "balinux"
	fmt.Println("structPointer: ", structPointer.name)
	fmt.Println("structPointer2: ", struct2pointer.name)

	// embedded struct
	embeddedData := embedded2{}
	embeddedData.grade = 2020
	embeddedData.name = "rio"
	embeddedData.age = 32

	fmt.Println("embedded struct")
	fmt.Println("name:", embeddedData.name)
	fmt.Println("age:", embeddedData.age)
	fmt.Println("age:", embeddedData.embedded1.age)
	fmt.Println("grade:", embeddedData.grade)

	// embedded struct dengan nama property yang sama
	samepropertyData := sameproperty2{}
	samepropertyData.name = "rio"
	samepropertyData.sameproperty1.age = 22
	samepropertyData.age = 33
	samepropertyData.grade = 2222

	fmt.Println("struct wish same property")
	fmt.Println("name:", samepropertyData.name)
	fmt.Println("age:", samepropertyData.sameproperty1.age)
	fmt.Println("age:", samepropertyData.age)
	fmt.Println("grade:", samepropertyData.grade)

	subData1 := substruct1{name: "rio", age: 22}
	subData2 := substruct2{substruct1: subData1, grade: 400}

	fmt.Println("substc section")
	fmt.Println("name:", subData2.name)
	fmt.Println("age:", subData2.age)
	fmt.Println("grade:", subData2.grade)

	// anonymous struct
	structAnonymData := struct {
		anonymStruct
		grade int
	}{}

	structAnonymData.anonymStruct = anonymStruct{"rio", 33}
	structAnonymData.grade = 22

	fmt.Println("anonymous struct")
	fmt.Println("name:", structAnonymData.anonymStruct.name)
	fmt.Println("age:", structAnonymData.anonymStruct.age)
	fmt.Println("grade:", structAnonymData.grade)

}

type anonymStruct struct {
	name string
	age  int
}

type substruct1 struct {
	name string
	age  int
}

type substruct2 struct {
	substruct1
	grade int
}

type sameproperty1 struct {
	name string
	age  int
}

type sameproperty2 struct {
	sameproperty1
	age   int
	grade int
}

type embedded1 struct {
	name string
	age  int
}

type embedded2 struct {
	grade int
	embedded1
}
