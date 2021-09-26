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
	fmt.Println("enter wight and lenght")
	var w, l float64
	fmt.Scanln(&w, &l)
	p := (w + l) * 2
	s := w * l
	fmt.Println("S=", s)
	fmt.Println("P=", p)
}
