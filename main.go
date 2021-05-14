package main

import (
	"fmt"
	"os"
	"weather_cli/api"
	"weather_cli/cli"
)

func main() {
	// cmd := []string{"cmd", "--city=\"London\"", "--unit=\"metric\""}
	city, unit, err := cli.Parse(os.Args)
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	var d api.WeatherData
	d, err = api.GetWeatherData(city, unit)
	if err != nil {
		fmt.Printf("a %v", err)
		os.Exit(1)
	}
	api.PrintWeather(d, city)
}
