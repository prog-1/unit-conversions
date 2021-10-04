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
	var a float64
	var b float64
	fmt.Print("Enter the length and the breadth of the rectangle: ")
	fmt.Scan(&a)
	fmt.Scan(&b)
	res1 := 2.0 * (a + b)
	res2 := a * b
	fmt.Printf("Perimeter: %.2f, square: %.2f", res1, res2)

}
