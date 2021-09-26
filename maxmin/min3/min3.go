package main

import (
	"fmt"
	"math"
)

func min(a, b float64) float64 {
	return (a + b - math.Abs(a-b)) / 2
}

func main() {
	fmt.Println("The program finds the maximum of three numbers max(a, b, c).")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Println("min= ", min(min(a, b), c))

}
