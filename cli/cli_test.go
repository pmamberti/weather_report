package cli_test

import (
	"testing"
	"weather_cli/api"
	"weather_cli/cli"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		args        []string
		errExpected bool
		name        string
		city        string
		unit        api.UnitSystem
	}{
		{
			name:        "No Flags, use defaults",
			args:        []string{"cmd"},
			errExpected: false,
			city:        "london",
			unit:        cli.UnitsMetric,
		},
		{
			name: "Valid parameters",
			args: []string{
				"cmd",
				`--city=Rome`,
				"--unit=standard",
			},
			errExpected: false,
			city:        "Rome",
			unit:        cli.UnitsStandard,
		},
		{
			name:        "Empty unit flag, default to `metric`",
			args:        []string{"cmd", `--city=Paris`, `--unit=""`},
			errExpected: true,
			city:        "Paris",
			unit:        cli.UnitsMetric,
		},
		{
			name:        "No unit flag, default to `metric`",
			args:        []string{"cmd", `--city=Paris`},
			errExpected: false,
			city:        "Paris",
			unit:        cli.UnitsMetric,
		},
		{
			name: "Empty Location flag",
			args: []string{
				"cmd",
				`--city=`,
				"--unit=\"metric\"",
			},
			errExpected: true,
			city:        "",
			unit:        cli.UnitsMetric,
		},
		{
			name: "Unit flag value not allowed",
			args: []string{
				"cmd",
				"--city=\"Las Palmas\"",
				`--unit="Hello"`,
			},
			errExpected: true,
			city:        "Las Palmas",
			unit:        cli.UnitsMetric,
		},
		{
			name: "Unit flag value not allowed",
			args: []string{
				"cmd",
				`--city="Moscow"`,
				`--unit=imperial`,
			},
			errExpected: true,
			city:        "Moscow",
			unit:        cli.UnitsMetric,
		},
	}

	for _, tc := range testCases {
		city, unit, err := cli.Parse(tc.args)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Errorf(
				"%v - Unexpected error received: %v, got %v instead",
				tc.name,
				tc.errExpected,
				errReceived,
			)
		}

		if !errReceived && !cmp.Equal(tc.city, city) {
			t.Errorf(
				"%v - want %v, got %v instead",
				tc.name,
				tc.city,
				city,
			)
		}

		if !errReceived && !cmp.Equal(tc.unit, unit) {
			t.Errorf(
				"%v - want %v, got %v instead",
				tc.name,
				tc.unit,
				unit,
			)
		}

	}
}
