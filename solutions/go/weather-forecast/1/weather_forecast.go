// Package weather is a forcast package where 
// you can get the current weather condition for a given city.
package weather


var (

    // CurrentCondition holds the current weather condition.
	CurrentCondition string

	// CurrentLocation holds the current location.
	CurrentLocation  string
)


// Forecast returns a string stating the current weather condition in a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
