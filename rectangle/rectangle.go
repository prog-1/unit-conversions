package main

import "fmt"

/*

Sample input/output:

The program prints the perimeter and the square of a rectangle given the rectangle sides.
Enter the length and the breadth of the rectangle:
4 5.5
Perimeter: 19
Square: 22

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	fmt.Println("Enter the resistance of two resistors:")
	var a, b float64
	fmt.Println("Enter the lenght an the breadth of the rectangle")
	fmt.Scanln(&a, &b)
	Perimeter := (a + b) * 2
	Square := a * b
	fmt.Println("Perimeter:", perimeter)
	fmt.Println("Square:", square)
}
