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
	fmt.Print("Enter the temperature in Celsius:")
	var t float64
	fmt.Scan(&t)
	t = (t * 9 / 5) + 32
	fmt.Println("The temperature in Fahrenheit:", t)
}
