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
	var r, per, sqr  float64
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	fmt.Println("Enter the radius: ")
	fmt.Scan(&r)
	per = 2 * r * math.Pi
	sqr = math.Pi * r * r
	fmt.Println("Perimetr: ", per)
	fmt.Println("square: ", sqr)

}
