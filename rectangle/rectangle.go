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
	var a,b float64
	fmt.Scanln(&a, &b)
	R := a + b
	G:= 1/a +1/b
	R3 := 1 / G
	fmt.Ptinln("Resistance when connected in series is", R, "resistance when connected in parallel is", R3)
}
