//+build integration

package weather_test

import "testing"
import "weather"

func TestGetWeatherDataIntegration(t *testing.T) {
	w, err := weather.GetWeatherData("London", weather.UnitsMetric)
	if err != nil {
		t.Fatal(err)
	}
	got :=w.Weather[0].Main
	if got == "" {
		t.Error("got empty main in weather data")
	}
}