// Package weather provides a way to tell the forecast on certain place.
package weather

// CurrentCondition provides the current weather condition of the CurrentLocation.
var CurrentCondition string
// CurrentLocation provides the current location where the CurrentCondition is being observed.
var CurrentLocation  string

// Forecast takes a city and a condition and returns a formatted string that tells what is the condition on that city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
