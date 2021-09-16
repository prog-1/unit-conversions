package main

import (
	"fmt"
	"math"
)

/*

Sample input/output:

The program prints the perimeter and the square of a circle given the radius.
Enter the radius: 8
Perimeter: 50.27, square: 201.06

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	fmt.Print("Enter the radius: ")
	var r float64
	fmt.Scanln(&r)
	fmt.Printf("Perimeter: %.2f, square: %.2f", 2*math.Pi*r, math.Pi*r*r)
	// ++ Your code here! ++
	//
	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
