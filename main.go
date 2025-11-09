package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Location struct {
	city    string
	state   string
	country string
}

type Coordinates struct {
	x float64
	y float64
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	loc, err := GetInput()

	if err == nil {
		log.Fatal("Error getting input")
	}

	GetCoordinates(loc)
}

func GetInput() (Location, error) {
	var loc Location

	fmt.Print("Enter a city name: ")
	fmt.Scanln(&loc.city)

	fmt.Print("Enter a city name: ")
	fmt.Scanln(&loc.state)

	fmt.Print("Enter a city name: ")
	fmt.Scanln(&loc.country)

	fmt.Println(loc.city, loc.state, loc.country)

	return loc, nil
}

func GetCoordinates(loc Location) Coordinates {
	var coords Coordinates

	url := "http://api.openweathermap.org/geo/1.0/direct?q="

	if loc.city != "" {
		url += loc.city + ","
	}
	if loc.state != "" {
		url += loc.state + ","
	} else {
		url += ","
	}
	if loc.country != "" {
		url += loc.country
	}

	apikey := os.Getenv("OPENWEATHER_API_KEY")
	if apikey == "" {
		log.Fatal("OpenWeather API key not set in .env")
	}

	url += "&limit=1&appid=" + apikey

	resp, err := http.Get(url)
	if err == nil {
		log.Fatal("Error fetching city coordinates")
	}
	defer resp.Body.Close()

	return coords
}
