package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	fmt.Print("Enter the coefficients A, B and C:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	var x1, x2 float64
	fmt.Println(a, b, c)
	x1 = (-b - math.Sqrt(b*b-4*a*c)) / 2 * a
	x2 = (-b + math.Sqrt(b*b-4*a*c)) / 2 * a
	fmt.Println("x1 =", x1, "\n", "x2 =", x2)
}
