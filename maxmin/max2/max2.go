package main

/* (maxmin/max2) a program that calculates the maximum of two numbers. You can use the following formula: max(a, b) = (a + b + |a - b|)/2. You will need to use math.Abs function to calculate an absolute value. Example:

The program finds the maximum of two numbers max(a, b).
Enter two numbers:
5 7
max: 7 */
import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of two numbers.")
	fmt.Println("Enter two numbers:")
	var a float64
	var b float64
	fmt.Scan(&a, &b)
	var max = math.Max(a, b)
	fmt.Printf("max: %v", max)
}
