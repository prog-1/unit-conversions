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
	fmt.Println("Enter Celsius temperature:")
	var t float64
	fmt.Scanln(&t)
	fmt.Println("result in Fahrenheit:", t/0.5556+32)
}
