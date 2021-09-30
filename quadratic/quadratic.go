package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	var A, B, C float64
	fmt.Println("Enter the coefficients A, B and C:")
	fmt.Scan(&A, &B, &C)
	D := B*B - 4*A*C
	x1 := (-B + math.Sqrt(D)) / (2 * A)
	x2 := (-B - math.Sqrt(D)) / (2 * A)
	fmt.Println("x1=", x1)
	fmt.Println("x2=", x2)

}
