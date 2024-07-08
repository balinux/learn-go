package main

import "fmt"

func main() {
	// deklarasi variabel map edngan formamt map[type_key]o_value
	var chicken map[string]int

	chicken = map[string]int{}

	// set varianel map
	chicken["jan"] = 400
	chicken["feb"] = 500

	// access map[type]type
	fmt.Println("januari: ", chicken["jan"])
	fmt.Println("februari: ", chicken["feb"])

	// inisialisasi nilai map
	// by defautl value map adalah nil, maka itu perlu di set
	data := map[string]string{
		"firstName": "rio",
		"lastname":  "putra",
	}
	println(data["firstName"], data["lastname"])

	// Iterasi itmem map dengan menggunakan for-renge
	donations := map[string]int{
		"rio":     300,
		"khaldi":  400,
		"sofwati": 500,
	}

	for key, value := range donations {
		fmt.Println(key, " \t:", value)
	}

	// deteksi keberadaan item dengan key tertentu

}
