package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program find solutions of a quadratic equation.")
	fmt.Println("Enter the coefficients A, B and C:")
	var A, B, C float64
	fmt.Scan(&A, &B, &C)
	x1 := (-B + math.Sqrt(B*B-4*A*C)) / (2 * A)
	x2 := (-B - math.Sqrt(B*B-4*A*C)) / (2 * A)
	fmt.Println("x1 =", x1)
	fmt.Println("x2 =", x2)
}
