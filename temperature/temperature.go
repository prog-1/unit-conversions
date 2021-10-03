package main

import "fmt"

func main() {
	fmt.Println("The program converts temperature from Celsius to Fahrenheit.")
	fmt.Println("Enter the temperature in Celsius:")
	var cels float64
	fmt.Scanln(&cels)
	fahrenh := cels*1.8 + 32
	fmt.Println("The temperature in Fahrenheit:", fahrenh)
}
