package main

import "fmt"

func main() {
	fmt.Println("Welcome to the temperature converter!")
	chosenPath := 'f'
	var tempFunction func(float64) float64
	if chosenPath == 'f' {
		tempFunction = celsiusToFahrenheit
	}
	temp := 0.0
	newTemp := tempFunction(temp)
	fmt.Printf("%v in Fahrenheint is %v\n", temp, newTemp)
}

func celsiusToFahrenheit(temp float64) float64 {
	return (temp * 9 / 5) + 32
}

func fahrenheitToCelsius(temp float64) float64 {
	return (temp - 32) * 5 / 9
}
