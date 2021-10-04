package main

import "fmt"

/*

Sample input/output:

The program prints the perimeter and the square of a rectangle given the rectangle sides.
Enter the length and the bredth of the rectangle:
4 5.5
Perimeter: 19
Square: 22

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	var a, b float64
	fmt.Println("Enter the lenght and the breadth of the rectangle")
	fmt.Scanln(&a, &b)
	fmt.Println("Perimeter:", (a+b)*2)
	fmt.Println("Square:", a*b)
}