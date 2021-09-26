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
	fmt.Println("enter C° temperature")
	var c float64
	fmt.Scanln(&c)
	f := c*1.8 + 32
	fmt.Println("F°=", f)
	// ++ Your code here! ++
}
