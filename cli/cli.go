package cli

import (
	"flag"
)

func Parse(args []string) (string, string) {

	city := flag.String("city", "London", "The city of which you want to know the current weather. Defaults to London")
	unit := flag.String("unit", "metric", "Choose your unit for the temperature (standard|metric|imperial). Defaults to metric.")

	flag.Parse()
	return *city, *unit

}
