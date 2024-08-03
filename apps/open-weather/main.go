package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	apiKey  = "28ea9963ad1cc250642c93f252e9cd22"
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
)

type WeatherResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func GetWeather(city string) (*WeatherResponse, error) {
	url := fmt.Sprintf("%s?q=%s&appid=%s&units=metric", baseURL, city, apiKey)

	// request data
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("error: %s", string(body))
	}

	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return nil, fmt.Errorf("failed to decode : %v", err)
	}
	return &weatherResponse, nil
}

func main() {
	city := "Denpasar"
	weater, err := GetWeather(city)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(weater.Name, weater.Main.Temp)
}
