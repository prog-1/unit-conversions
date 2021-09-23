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
	var rad float64
	fmt.Print("Enter the radius:")
	fmt.Scan(&rad)
	var per float64
	per = rad * 2 * math.Pi
	fmt.Println("Perimeter:", per)
	var sq float64
	sq = rad * rad * math.Pi
	fmt.Println("Square:", sq)

}
