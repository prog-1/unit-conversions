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
	fmt.Print("Enter the length and the breadth of the rectangle:")
	var a, b float64
	fmt.Scan(&a, &b)
	p := (a + b) * 2
	s := a * b
	fmt.Println("Perimeter:", p)
	fmt.Println("Square:", s)
}
