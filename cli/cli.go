package cli

import (
	"errors"
	"flag"
)

<<<<<<< HEAD
var c, u *string

func init() {
	c = flag.String(
		"city",
		"london",
		"enter city - required",
	)
	u = flag.String(
		"unit",
		"metric",
		"pick one between metric and standard",
	)
}

=======
>>>>>>> bdbf5d3eb4e9bf0352de31e4158ea3ab02ecfe0f
func Allowed(list []string, u string) bool {
	for _, s := range list {
		if s == u {
			return true
		}
	}
	return false
}

func Parse(cmd []string) ([]string, error) {
	var err error
	arguments := flag.NewFlagSet("args", flag.ContinueOnError)
	c := arguments.String("city", "london", "Required. Target city")
	u := arguments.String("unit", "metric", "Optional. Temperature unit: Standard, Metric (default) or Imperial.")
	arguments.Parse(cmd[1:])

	if *c == "" {
		err = errors.New("city cannot be empty - default to london")
		arguments.PrintDefaults()
		// os.Exit(1)
	}

	if !Allowed([]string{"metric", "standard", "imperial"}, *u) {
		*u = "metric"
		err = errors.New("unit must be metric standard or imperial")
	}

	return []string{*c, *u}, err
}
