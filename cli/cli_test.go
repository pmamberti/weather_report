package cli_test

import (
	"os"
	"testing"
	"weather_cli/cli"
)

func TestParse(t *testing.T) {
	t.Parallel()
	locwant := "London"
	unitwant := "metric"
	locgot, unitgot := cli.Parse(os.Args)

	if locgot != locwant || unitgot != unitwant {
		t.Errorf("want %v - %v, got %v - %v", locwant, unitwant, locgot, unitgot)
	}
}
