package main

import "fmt"

func main() {
	fmt.Println("The program converts temperature from Celsius to Fahrenheit.")
	var temperatureC float64
	fmt.Print("Enter the temperature in Celsius:")
	fmt.Scan(&temperatureC)
	fmt.Println("The temperature in Fahrenheit:", temperatureC*9/5+32)

}
