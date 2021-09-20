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
	fmt.Print("Enter the length and the breadth of the rectangle: ")
	var x, y float64
	fmt.Scan(&x, &y)
	Perimeter := (x + y) * 2
	Square := x * y
	fmt.Println("Perimeter:", Perimeter)
	fmt.Println("Square:", Square)
}
