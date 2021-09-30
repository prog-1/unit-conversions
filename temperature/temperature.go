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
	var temp_C float64
	fmt.Print("Enter the temperature in Celsius:")
	fmt.Scan(&temp_C)
	var temp_F float64
	temp_F = temp_C*9/5 + 32
	fmt.Println("The temperature in Fahrenheit:", temp_F)
}
