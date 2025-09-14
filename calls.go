package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Name string `json:"name"`
	Weather	[]Weather `json:"weather"`
	Main WeatherMain `json:"main"`
	Wind WeatherWind `json:"wind"`
}

type Weather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type WeatherMain struct {
	Temp float32 `json:"temp"`
	Feels_like float32 `json:"feels_like"`
	Temp_min float32 `json:"temp_min"`
	Temp_max float32 `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"`
	Sea_level int `json:"sea_level"`
	Grnd_level int `json:"grnd_level"`
}

type WeatherWind struct {
	Speed float32 `json:"speed"`
	Deg int `json:"deg"`
}

func kelvinToCelsius(kelvin float32) float32 {
	return kelvin - 273.15
}

func FetchWeatherData() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env")
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	lat := 51.5072  // example: London
	lon := -0.1276  // example: London

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v",
		lat,
		lon,
		apiKey,
	)
	
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObj WeatherResponse
	json.Unmarshal(data, &responseObj)

	fmt.Printf("Name: %v\n", responseObj.Name)
	fmt.Printf("Weather: %v\n", responseObj.Weather[0].Main)
	fmt.Printf("Main Temp: %v\n", kelvinToCelsius(responseObj.Main.Temp))
	fmt.Printf("Main Feels Like: %v\n", kelvinToCelsius(responseObj.Main.Feels_like))
	fmt.Printf("Main Temp Min: %v\n", kelvinToCelsius(responseObj.Main.Temp_min))
	fmt.Printf("Main Temp Max: %v\n", kelvinToCelsius(responseObj.Main.Temp_max))
	fmt.Printf("Main Pressure: %v\n", responseObj.Main.Pressure)
	fmt.Printf("Main Humidity: %v\n", responseObj.Main.Humidity)
	fmt.Printf("Main Sea Level: %v\n", responseObj.Main.Sea_level)
	fmt.Printf("Main Grnd Level: %v\n", responseObj.Main.Grnd_level)
	fmt.Printf("Wind Speed: %v\n", responseObj.Wind.Speed)

} 