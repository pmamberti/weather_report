package cli_test

import (
	"testing"
	"weather_cli/cli"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		args        []string
		errExpected bool
		name        string
		want        []string
	}{
		{
			name:        "No Flags, use defaults.",
			args:        []string{"cmd"},
			errExpected: false,
			want:        []string{"london", "metric"},
		},
		{
			name: "Valid parameters",
			args: []string{
				"cmd",
				`--city=Rome`,
				"--unit=standard",
			},
			errExpected: false,
			want:        []string{"Rome", "standard"},
		},
		{
			name:        "Empty unit flag, default to `metric`",
			args:        []string{"cmd", `--city=Paris`, `--unit=""`},
			errExpected: true,
			want:        []string{"Paris", "metric"},
		},
		{
			name:        "No unit flag, default to `metric`",
			args:        []string{"cmd", `--city=Paris`},
			errExpected: false,
			want:        []string{"Paris", "metric"},
		},
		{
			name: "Empty Location flag",
			args: []string{
				"cmd",
				`--city=`,
				"--unit=\"metric\"",
			},
			errExpected: true,
			want:        []string{"", "metric"},
		},
		{
			name: "Unit flag value not allowed",
			args: []string{
				"cmd",
				"--city=\"Las Palmas\"",
				`--unit="Hello"`,
			},
			errExpected: true,
			want:        []string{"Las Palmas", "metric"},
		},
	}

	for _, tc := range testCases {
		got, err := cli.Parse(tc.args)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Errorf(
				"%v - Unexpected error received: %v, got %v instead",
				tc.name,
				tc.errExpected,
				errReceived,
			)
		}

		if !errReceived && !cmp.Equal(tc.want, got) {
			t.Errorf(
				"%v - want %v, got %v instead",
				tc.name,
				tc.want,
				got,
			)
		}
	}
}
