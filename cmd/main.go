package main

import (
	"os"
	weather "weather_cli"
)

func main() {
	weather.Run(os.Args)
}
