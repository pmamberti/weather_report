package weather

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type owmAPIResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	}
}

type WeatherData struct{
	Summary string
	Description string
	Temp float64
}

type UnitSystem int

var unitSystemName = map[UnitSystem]string{
	0: "metric",
	1: "standard",
	2: "imperial",
}

type Client struct {
	APIKey string
	BaseURL string
	HTTPClient *http.Client
}

func NewClient(APIKey string) (Client, error) {
	if APIKey == "" {
		return Client{}, errors.New("empty API key")
	}
	return Client{
		APIKey: APIKey,
		BaseURL: "https://api.openweathermap.org",
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}, nil
}

func (c Client) GetWeatherData(location string, unit UnitSystem) (WeatherData, error) {
	url := fmt.Sprintf(
		"%s/data/2.5/weather?units=%v&q=%v&appid=%v",
		c.BaseURL,
		unitSystemName[UnitSystem(unit)],
		url.PathEscape(location),
		c.APIKey,
	)

	res, err := c.HTTPClient.Get(url)

	if err != nil {
		log.Fatalf("Cannot complete GET request: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		// TODO: check how to get the message and make this more informative
		// e.g. {"cod":"404","message":"city not found"}%
		log.Fatalf("Unexpected Status: %v", res.StatusCode)
	}

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Cannot read data from request body: %v", err)
	}
	var data owmAPIResponse
	err = json.Unmarshal(resData, &data)
	if err != nil {
		log.Fatalf("Unable to Marshal data %v", err)
	}

	return WeatherData{
		Summary: data.Weather[0].Main,
		Description: data.Weather[0].Description,
		Temp: data.Main.Temp,
	}, nil
}

func PrintWeather(d WeatherData, location string) {

	fmt.Printf(
		"In %v, the weather is %v, with %v. \nThe temperature is %v degrees.",
		location,
		d.Summary,
		d.Description,
		d.Temp,
	)
}


const (
	UnitsMetric = iota
	UnitsStandard
	UnitsImperial
)

func convertUnits(s string) (UnitSystem, error) {
	switch s {
	case "metric":
		return UnitsMetric, nil
	case "standard":
		return UnitsStandard, nil
	case "imperial":
		return UnitsImperial, nil
	default:
		return 0, fmt.Errorf("value not allowed %q must be one of metric standard or imperial", s)
	}

}

func Parse(cmd []string) (city string, unit UnitSystem, err error) {
	arguments := flag.NewFlagSet("args", flag.ExitOnError)
	c := arguments.String("city", "london", "Required. Target city")
	u := arguments.String("unit", "metric", "Optional. Temperature unit: Standard, Metric (default) or Imperial.")

	if len(cmd) < 2 {
		arguments.Parse([]string{})
	} else {
		arguments.Parse(cmd[1:])
	}

	unit, err = convertUnits(*u)
	if err != nil {
		return "", 0, err
	}

	if *c == "" {
		err = errors.New("city cannot be empty - default to london")
		arguments.PrintDefaults()
	}

	return *c, unit, err
}

func RunCLI(args []string, w io.Writer) {
	city, unit, err := Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	api_key := os.Getenv("OWM_KEY")
	if api_key == "" {
		fmt.Fprintln(os.Stderr, "Please set OWM_KEY environment variable")
		os.Exit(1)
	}
	client, err := NewClient(api_key)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var d WeatherData
	d, err = client.GetWeatherData(city, unit)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	PrintWeather(d, city)
}