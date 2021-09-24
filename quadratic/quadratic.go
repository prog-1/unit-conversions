package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	fmt.Println("Enter the coefficients a, c and b:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	D := b*b - (4 * a * c)
	x1 := (-b + math.Sqrt(D)) / 2
	x2 := (-b - math.Sqrt(D)) / 2
	fmt.Println("x1 =", x1)
	fmt.Println("x2 =", x2)
}
