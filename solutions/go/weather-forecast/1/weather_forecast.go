// Package weather provides tools to provide
// current weather conditions.
package weather

// CurrentCondition variable define weather conditions.
var CurrentCondition string

// CurrentLocation variable define current location.
var CurrentLocation string

// Forecast function return a combined string with the current weather of the provided location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
