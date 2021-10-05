package main

import (
	"fmt"
)

/*

Sample input/output:

The program prints the perimeter and the square of a circle given the radius.
Enter the radius:
8
Perimeter: 50.26548245743669
Square: 201.06192982974676

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	var radius float64
	fmt.Println("Enter the radius:")
	fmt.Scanln(&radius)
	perimeter := radius * 2 * math.Pi 
	var square float64
	square := radius * radius * math.Pi
	fmt.Println("Perimeter", perimeter)
	fmt.Println("Square", square)
}