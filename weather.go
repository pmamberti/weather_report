package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type WeatherData struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	}
}

func weather() {
	api_key := os.Getenv("OWM_KEY")
	if api_key == "" {
		log.Fatalf("No API key found in env")
	}
	location := strings.Join(os.Args[1:], " ")
	if cmp.Equal(os.Args[1:], []string{}) {
		log.Fatalf("No city specified, exiting...")
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?units=metric&q=%v&appid=%v", url.PathEscape(location), api_key)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Cannot complete GET request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected Status: %v", res.StatusCode)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Cannot read data from body: %v", err)
	}

	var j WeatherData
	err = json.Unmarshal(data, &j)
	if err != nil {
		log.Fatalf("Unable to Marshal data %v", err)
	}

	desc := j.Weather[0].Main
	temp := j.Main.Temp
	fmt.Printf("In %v, the weather is \"%v\" and the temperature is %v degrees (Celsius, obviously!)", location, desc, temp)
}
