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
	fmt.Println("Enter the temperature in Celsius:")
	var cls float64
	fmt.Scan(&cls)
	var frngt float64 = 1.8*cls + 32
	fmt.Printf("The temperature in Fahrenheit: %v", frngt)
}
