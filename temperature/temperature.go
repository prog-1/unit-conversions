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
	// ++ Your code here! ++
	var t float64
	fmt.Println("Enter the temperature in Celsius:")
	fmt.Scan(&t)
	fmt.Println("The temperature in Fahrenheit:", t*1.8+32)
}
