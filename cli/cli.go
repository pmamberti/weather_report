package cli

import (
	"flag"
)

type Data struct {
	Location string
	Unit     string
}

func Parse(args []string) Data {
	var data = Data{}

	city := flag.String("city", "London", "The city of which you want to know the current weather. Defaults to London")
	unit := flag.String("unit", "metric", "Choose your unit for the temperature (metric|kelvin|meh). Defaults to metric.")
	flag.Parse()

	data.Location = *city
	data.Unit = *unit
	return data
}
