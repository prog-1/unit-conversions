package main

import "fmt"

/*
-
Sample input/output:

The program converts temperature from Celsius to Fahrenheit.
Enter the temperature in Celsius: 100
The temperature in Fahrenheit: 212.00

*/
func main() {
	fmt.Println("The program converts temperature from Celsius to Fahrenheit.")
	fmt.Print("Enter the temperature in Celsius: ")
	var cel float64
	fmt.Scan(&cel)
	fmt.Printf("The temperature in Fahrenheit: %.2f", (cel*9/5)+32)
}
