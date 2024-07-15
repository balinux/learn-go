package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "https://www.yhotie.com/menambah-domain-ke-vercel?name=rio&age=22"
	u, e := url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url:%s\n", urlString)
	fmt.Printf("protocol:%s\n", u.Scheme)
	fmt.Printf("host:%s\n", u.Host)
	fmt.Printf("path:%s\n", u.Path)

	name := u.Query()["name"][0]
	age := u.Query()["age"][0]

	fmt.Printf("name:%s\n", name)

	fmt.Printf("age:%s\n", age)

}
