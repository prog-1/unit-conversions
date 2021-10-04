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
	fmt.Print("Enter the radius: ")
	var r float64
	fmt.Scan(&r)
	fmt.Printf("Perimeter: %.2f, square: %.2f", 2*r*math.Pi, math.Pi*r*r)
	// ++ Your code here! ++
	//
	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
