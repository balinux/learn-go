package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	user, err := fetchUsers()
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("User: %+v\n", user) // Access and print the first name and last name
	fmt.Printf("First Name: %s\n", user.Data.FirstName)
	fmt.Printf("Last Name: %s\n", user.Data.LastName)
}

type Data struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type Support struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}

type Response struct {
	Data    Data    `json:"data"`
	Support Support `json:"support"`
}

func fetchUsers() (*Response, error) {
	baseUrl := "https://reqres.in/api/users/2"

	// make http get request
	resp, err := http.Get(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user data: %v", err)
	}
	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// unmarshal json ke object struct
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	return &response, nil
}
