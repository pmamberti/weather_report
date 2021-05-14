package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
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


func TestGetWeatherData(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("testdata/owm.json")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}))
	APIKey := "dummy"
	client, err := weather.NewClient(APIKey)
	if err != nil {
		t.Fatal(err)
	}
	client.BaseURL = ts.URL
	client.HTTPClient = ts.Client()
	got, err := client.GetWeatherData("London", weather.UnitsMetric)
	if err != nil {
		t.Fatal(err)
	}
	want := weather.WeatherData{
		Summary: "Clouds",
		Description: "overcast clouds",
		Temp: 12.04,
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}