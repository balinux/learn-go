package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("asset"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"Name":    "rio",
			"Message": "belajar tempalte golang web",
		}
		t, err := template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("starting webs erver at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
