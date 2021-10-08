package main

import "fmt"

func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	var l, b float64
	fmt.Println("Enter the length and the breadth of the rectangle:")
	fmt.Scanln(&l, &b)
	fmt.Println("Perimetr:", (l+b)*2)
	fmt.Println("Square:", l*b)
}
