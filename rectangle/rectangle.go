package main

import "fmt"

/*

Sample input/output:

The program prints the perimeter and the square of a rectangle given the rectangle sides.
Enter the length and the breadth of the rectangle:
4 5.5
Perimeter: 19
Square: 22
//
*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	// ++ Your code here! ++
	fmt.Println("Enter the length and the breadth of the rectangle:")
	var aside float64
	var bside float64
	fmt.Scan(&aside, &bside)
	var prm = 2 * (aside + bside)
	var sqr = aside * bside
	fmt.Printf("Perimeter: %v\n", prm)
	fmt.Printf("Square: %v", sqr)
}
