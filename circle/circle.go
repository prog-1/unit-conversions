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
	// ++ Your code here! ++
	//
	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
	var r float64
	const π = math.Pi
	fmt.Println("Enter the radius: ")
	fmt.Scan(&r)
	fmt.Println("Perimeter:", 2*π*r)
	fmt.Println("Square:", π*r*r)
}
