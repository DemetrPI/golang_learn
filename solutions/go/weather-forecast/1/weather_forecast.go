// Package weather provides
// weather conditions in current location.
package weather

// CurrentCondition is a variable to save current condition state.
var CurrentCondition string

// CurrentLocation is a variable to set current location.
var CurrentLocation string

//Forecast function returns weather condition for current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
