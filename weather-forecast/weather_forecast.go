// Package weather This is the package for whether it will return all the endpoints and other things.
package weather

// CurrentCondition Variables are used to store data in programs.
var CurrentCondition string

// CurrentLocation Variables are used to store data in programs.
var CurrentLocation string

// Forecast This the funtion that calculates the values for current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
