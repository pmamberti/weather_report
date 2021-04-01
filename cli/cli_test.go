package cli_test

import (
	"fmt"
	"testing"
	"weather_cli/cli"
)

func TestParseInput(t *testing.T) {
	t.Parallel()
	want := cli.Data{Location: "las palmas", Unit: "kelvin"}
	testInput := fmt.Sprintf("cmd --city=%v --unit=%v", want.Location, want.Unit)

	got := cli.ParseInput(testInput)

	if got != want {
		t.Errorf("want %v, got %v", want, got)
	}
}
