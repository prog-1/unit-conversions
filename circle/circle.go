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
	fmt.Print("Enter the radius:")
	var r float64
	fmt.Scan(&r)
	p := 2 * math.Pi * r
	s := math.Pi * r * r
	fmt.Println("Perimeter:", p)
	fmt.Println("Square:", s)
}
