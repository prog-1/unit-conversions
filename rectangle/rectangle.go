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
	fmt.Println("Enter the lenght and the breadth of the rectangle:")
	var x, y float64
	fmt.Scanln(&x, &y)
	p := (x + y) * 2
	s := x * y
	fmt.Println("Perimeter:", p)
	fmt.Println("Square:", s)
}
