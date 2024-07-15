package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	// casting ke byte
	jsonData := []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("username: ", data.FullName)
	fmt.Println("age: ", data.Age)

	// Decode JSON Ke map[string]interface{} & interface{}
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("decode json ke map")
	fmt.Println("user", data1["Name"])
	fmt.Println("age", data1["Age"])

	// decode JSON ke interface
	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	// casting lebih dahulu sebelum akses data
	var decodeData = data2.(map[string]interface{})
	fmt.Println("decode to interface")
	fmt.Println("user", decodeData["Name"])
	fmt.Println("age", decodeData["Age"])
}
