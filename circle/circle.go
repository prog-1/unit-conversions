package main

import (
	"fmt"
	"math"
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
	fmt.Print("Enter the radius:")
	fmt.Scan(&radius)
	var perimeter float64
	perimeter = radius * 2 * math.Pi
	fmt.Println("Enter the perimeter:", perimeter)
	var square float64
	square = radius * radius * math.Pi
	fmt.Println("Enter the square:", square)
	//
	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
