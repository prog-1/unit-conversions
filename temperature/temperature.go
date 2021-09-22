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
	var T float64
	fmt.Println("Enter the temperature in Celsius:")
	fmt.Scan(&T)
	fmt.Println("The temperature in Fahrenheit:", T*1.8+32)
}
