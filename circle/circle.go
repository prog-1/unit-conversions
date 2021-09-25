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
	// ++ Your code here! ++
	fmt.Println("Enter the radius:")
	var rds float64
	fmt.Scan(&rds)
	var Pi float64 = 3.14159265358979323846264338327950288419716939937510582097494459
	var prm = 2 * (Pi * rds)
	var sqr = Pi * (rds * rds)
	fmt.Printf("Perimeter: %v\n", prm)
	fmt.Printf("Square: %v", sqr)

	//
	// Hint: Use math.Pi constant (https://pkg.go.dev/math#pkg-constants) to obtain π value.
}
