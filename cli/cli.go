package cli

import (
	"errors"
	"flag"
	"fmt"
	"weather_cli/api"
)

const (
	UnitsMetric = iota
	UnitsStandard
	UnitsImperial
)

func convertUnits(s string) (api.UnitSystem, error) {
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

func Parse(cmd []string) (city string, unit api.UnitSystem, err error) {
	arguments := flag.NewFlagSet("args", flag.ContinueOnError)
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
