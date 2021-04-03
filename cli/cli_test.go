package cli_test

import (
	"testing"
)

func TestParse(t *testing.T) {

	testCases := []struct {
		args        []string
		errExpected bool
		name        string
		want        []string
	}{
		{args: []string{"cmd", "--city=\"London\"", "--unit=\"metric\""}, errExpected: false, want: []string{"London", "metric"}, name: "Default Parameters"},
		{args: []string{"cmd", "--city=\"Rome\"", "--unit=\"standard\""}, errExpected: false, want: []string{"London", "metric"}, name: "Different, still valid, parameters"},
		// {args: &[]string{"cmd", "--unit=\"metric\""}, errExpected: true, want: []string{"London", "metric"}, name: "Empty Location"},
		// {args: &[]string{"cmd", "--city=\"London\"", "--unit=\"\""}, errExpected: false, want: []string{"London", "metricc"}, name: "Empty Unit, default to `metric`"},
	}
}
