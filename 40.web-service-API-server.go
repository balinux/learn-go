package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)
	fmt.Println("starting API server at port 8000")
	http.ListenAndServe(":8000", nil)
}

// menampilkan daftar users
func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		r, e := json.Marshal(data)

		if e != nil {
			fmt.Println(e.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(r)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

// mengambil 1 user saja
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		id := r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				w.Write(result)
				return
			}
		}
		http.Error(w, "user not found", http.StatusNotFound)
	}
	http.Error(w, "", http.StatusBadRequest)
}
