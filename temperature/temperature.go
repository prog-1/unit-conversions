package main

import "fmt"

/*

Sample input/output:

The program converts temperature from Celsius to Fahrenheit.
Enter the temperature in Celsius:
100
The temperature in Fahrenheit: 212

*/
func main() {
	fmt.Println("The program converts temperature from Celsius to Fahrenheit.")
	var temperature float64
	fmt.Println("Celsius")
	fmt.Scan(&temperature)
	fmt.Println("Farenheit:", temperature*1.8+32)

}
