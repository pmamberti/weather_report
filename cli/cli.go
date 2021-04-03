package cli

import (
	"errors"
	"flag"
	"os"
)

var c, u *string

func init() {
	c = flag.String("city", "london", "enter city - required")
	u = flag.String(
		"unit",
		"metric",
		"pick one between metric and standard",
	)
}

func Allowed(list []string, u string) bool {
	for _, s := range list {
		if s == u {
			return true
		}
	}
	return false
}

func Parse(cmd []string) ([]string, error) {
	os.Args = cmd
	allowedUnits := []string{"metric", "standard", "imperial"}
	var err error
	flag.Parse()
	if *c == "" {
		err = errors.New("city not provided, you must specify one")
	}
	if !Allowed(allowedUnits, *u) || *u == "" {
		err = errors.New(
			"measure unit not allowed, defaulting to metric ",
		)
		*u = "metric"
	}

	return []string{*c, *u}, err
}
