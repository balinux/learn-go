package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "rio jp"
	encodeString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodeString)

	decodeByte, _ := base64.StdEncoding.DecodeString(encodeString)
	println("decodeByte:", decodeByte)
	decodeString := string(decodeByte)
	fmt.Println(decodeString)

	//penerapan fungsi encode decode
	dataName := "muhammad khalid"
	var encodedName = make([]byte, base64.StdEncoding.EncodedLen(len(dataName)))
	base64.StdEncoding.Encode(encodedName, []byte(dataName))
	encodeStringName := string(encodedName)
	fmt.Println(encodeStringName)

	decodedName := make([]byte, base64.RawStdEncoding.DecodedLen(len(encodedName)))
	_, err := base64.StdEncoding.Decode(decodedName, encodedName)
	if err != nil {
		fmt.Println(err.Error())
	}
	decodeStringName := string(decodedName)
	fmt.Println(decodeStringName)
}
