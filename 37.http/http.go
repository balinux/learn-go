package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hallo saya rio")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello web server pertamaku pake golang")
	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
