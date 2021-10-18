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
Perimeter: c
Square: 201.06192982974676

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	var radius float64
	fmt.Println("radius:")
	fmt.Scanln(radius)
	fmt.Println("Peremiter:", 2*math.Pi*radius)
	fmt.Println("Square:", math.Pi*(radius*radius))

	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
