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
	fmt.Println("Enter the temperature in Celsius: ")
	var x float64
	fmt.Scan(&x)
	x = x*9/5 + 32
	fmt.Println("The temperature in Fahrenheit:", x)

}
