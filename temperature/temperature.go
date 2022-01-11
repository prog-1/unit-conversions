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
	var a float64
	fmt.Print("Enter the temperature in Celsius:")
	fmt.Scan(&a)
	fmt.Println("The temperature in Fahrenheit: ", ((a * 9 / 5) + 32))
}
