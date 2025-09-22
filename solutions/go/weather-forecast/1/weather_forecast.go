// Package weather represents a package which describes weather.
package weather

// CurrentCondition represents variable string,
// consists of weather condition.
var CurrentCondition string

// CurrentLocation represents geo location of 
// target weather.
var CurrentLocation string

// Forecast function returns current weather in target Location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
