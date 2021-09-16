package main

import "fmt"

/*

Sample input/output:

The program prints the perimeter and the square of a rectangle given the rectangle sides.
Enter the length and the bredth of the rectangle: 4 5.5
Perimeter: 19.00, square: 22.00

*/
func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	fmt.Print("Enter the length and the bredth of the rectangle: ")
	var l, b float64
	fmt.Scanln(&l, &b)
	fmt.Printf("Perimter: %.2f, square: %.2f\n", (l+b)*2, l*b)
	// ++ Your code here! ++
}
