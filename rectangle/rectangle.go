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
	var l, b float64
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	fmt.Println("Enter the length and the breadth of the rectangle: ")
	fmt.Scan(&l, &b)
	p := (l + b) *2 
	s := l * b 
	fmt.Println("Perimetr: ", p)
	fmt.Println("square: ", s)

}
