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
	fmt.Println("Enter radius:")
	var r float64
	fmt.Scan(&r)
	const p = math.Pi
	fmt.Println("Circle perimeter:", 2*p*r)
	fmt.Println("Circle square", p*r*r)

}
