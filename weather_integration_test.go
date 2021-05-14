//+build integration

package weather_test

import (
	"fmt"
	"os"
	"testing"
	"weather"
)

func TestGetWeatherDataIntegration(t *testing.T) {
	APIKey := os.Getenv("OWM_KEY")
	if APIKey == "" {
		fmt.Fprintln(os.Stderr, "Please set OWM_KEY environment variable")
		os.Exit(1)
	}
	client, err := weather.NewClient(APIKey)
	if err != nil {
		t.Fatal(err)
	}
	w, err := client.GetWeatherData("London", weather.UnitsMetric)
	if err != nil {
		t.Fatal(err)
	}
	got :=w.Weather[0].Main
	if got == "" {
		t.Error("got empty main in weather data")
	}
}