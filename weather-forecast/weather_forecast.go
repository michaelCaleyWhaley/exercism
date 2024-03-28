// Package weather provides a function to get the weather forecast at a given city.
package weather

// CurrentCondition is a description of the current weather conditionins in a given city.
var CurrentCondition string

// CurrentLocation respresents the selected location of the forecast.
var CurrentLocation string

// Forecast returns a string weather forecast containin the location and current conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
