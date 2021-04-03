package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type WeatherData struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	}
}

func GetWeather(location, unit string, err error) {
	api_key := os.Getenv("OWM_KEY")
	if api_key == "" {
		log.Fatalf("No API key found in env")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?units=%v&q=%v&appid=%v", unit, url.PathEscape(location), api_key)
	res, err := http.Get(url)
	fmt.Println(url)
	if err != nil {
		log.Fatalf("Cannot complete GET request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		// TODO: check how to get the message and make this more informative
		// e.g. {"cod":"404","message":"city not found"}%
		log.Fatalf("Unexpected Status: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Cannot read data from request body: %v", err)
	}

	var j WeatherData
	err = json.Unmarshal(data, &j)
	if err != nil {
		log.Fatalf("Unable to Marshal data %v", err)
	}

	fmt.Printf("In %v, the weather is %v, with %v. \nThe temperature is %v degrees.", location, j.Weather[0].Main, j.Weather[0].Description, j.Main.Temp)
}
