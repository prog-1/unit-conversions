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
	fmt.Println("enter radius")
	var r float64
	fmt.Scanln(&r)
	p := r * 2 * math.Pi
	s := r * r * math.Pi
	fmt.Println("S=", s)
	fmt.Println("P=", p)
	// ++ Your code here! ++
	//
	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
