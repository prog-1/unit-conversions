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
	var l, w float64
	fmt.Println("Enter the length and the bredth of the rectangle:")
	fmt.Scanln(&l, &w)
	fmt.Println("Perimeter:", (l+w)*2)
	fmt.Println("Square:", l*w)
}
