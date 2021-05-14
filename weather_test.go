package weather_test

import (
	"testing"
	"weather"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		args        []string
		errExpected bool
		name        string
		city        string
		unit        weather.UnitSystem
	}{
		{
			name:        "No Flags, use defaults",
			args:        []string{"cmd"},
			errExpected: false,
			city:        "london",
			unit:        weather.UnitsMetric,
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
			unit:        weather.UnitsStandard,
		},
		{
			name:        "Empty unit flag, default to `metric`",
			args:        []string{"cmd", `--city=Paris`, `--unit=""`},
			errExpected: true,
			city:        "Paris",
			unit:        weather.UnitsMetric,
		},
		{
			name:        "No unit flag, default to `metric`",
			args:        []string{"cmd", `--city=Paris`},
			errExpected: false,
			city:        "Paris",
			unit:        weather.UnitsMetric,
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
			unit:        weather.UnitsMetric,
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
			unit:        weather.UnitsMetric,
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
			unit:        weather.UnitsMetric,
		},
	}

	for _, tc := range testCases {
		city, unit, err := weather.Parse(tc.args)
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
