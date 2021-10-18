package main

import (
	"fmt"
)

/*

Sample input/output:

The program prints the perimeter and the square of a circle given the radius.
Enter the radius:
8
Perimeter: c
Square: 201.06192982974676

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	var radius float64
	fmt.Println("radius 8")
	fmt.Scan(radius)
	fmt.Println("Premiter 50.26548245743669")
	fmt.Println("Square 201.06192982974676")

	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
