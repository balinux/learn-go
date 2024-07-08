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

	// menghapus item map
	map1 := map[string]int{
		"rio":   40,
		"rian":  50,
		"linda": 60,
	}

	fmt.Println(map1)

	delete(map1, "linda")

	fmt.Println(map1)

	// deteksi keberadaan item dengan key tertentu
	data2 := map[string]int{
		"rio":   33,
		"rian":  44,
		"linda": 11,
	}
	value, isExist := data2["daos"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("data daos tidak tersedia")
	}

	// kombinasi slice dan map
	data3 := []map[string]string{
		{"name": "rio", "gender": "male"},
		{"name": "linda", "gender": "female"},
		{"name": "rian", "gender": "male"},
	}

	for _, data := range data3 {
		fmt.Println(data["name"], data["gender"])
	}
}
