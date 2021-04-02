package main

import (
	"os"
	api "weather_cli"
	"weather_cli/cli"
)

func main() {
	api.GetWeather(cli.Parse(os.Args))
}
