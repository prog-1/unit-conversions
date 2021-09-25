package main

import (
	"fmt"
	"math"
)

/*Given coefficients A, B, and C of a quadratic equation Ax^2 + Bx + C = 0, write a program that prints the equation solutions x1 and x2.

You will need to use math.Sqrt function to calculate a square root.

Example:

The program finds solutions of a quadratic equation.
Enter the coefficients A, B and C:
1 3 -10
x1 = 2
x2 = -5 */
func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	fmt.Println("Enter the coefficients A, B and C:")
	var A float64
	var B float64
	var C float64
	fmt.Scan(&A, &B, &C)
	var D float64 = B*B - 4*A*C
	if D > 0 {
		var x1 float64 = (-B + math.Sqrt(D)) / (2 * A)
		var x2 float64 = (-B - math.Sqrt(D)) / (2 * A)
		fmt.Printf("x1: %v\n", x1)
		fmt.Printf("x2: %v", x2)
	}
	if D == 0 {
		var x float64 = -B / (2 * A)
		fmt.Printf("x: %v", x)
	}
	if D < 0 {
		fmt.Println("No answer.")
	}
}
