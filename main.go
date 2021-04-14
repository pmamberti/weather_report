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
	fmt.Println(city, unit, err)
	api.GetWeather(city, unit)
}
