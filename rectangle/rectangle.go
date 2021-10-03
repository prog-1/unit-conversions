package main

import "fmt"

func main() {
	fmt.Println("The program prints the perimeter and the square of a rectangle given the rectangle sides.")
	fmt.Println("Enter the lenght and the breadth of the rectangle:")
	var l, b float64
	fmt.Scanln(&l, &b)
	p := (l + b) * 2
	s := l * b
	fmt.Println("Perimeter:", p)
	fmt.Println("Square:", s)
}
