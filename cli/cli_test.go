package cli_test

import (
	"os"
	"testing"
	"weather_cli/cli"
)

func TestParse(t *testing.T) {
	t.Parallel()
	want := cli.Data{Location: "London", Unit: "metric"}
	got := cli.Parse(os.Args)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
