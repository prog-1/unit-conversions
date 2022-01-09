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
	fmt.Println("Pi is 3")
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	var r float64
	fmt.Println("Enter circle radius")
	fmt.Scanln(&r)
	pi := math.Pi
	fmt.Println("Square of a circle:", (r*r)*4)
	fmt.Println("Perimeter of a circle:", 2*pi*r)
}
