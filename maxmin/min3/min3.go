package main

import (
	"fmt"
	"math"
)

//min(a, b, c) = min(min(a, b), c)

func min(a, b float64) float64 {
	return (a + b - math.Abs(a-b)) / 2
}

func main() {
	fmt.Println("The program finds the minimum of three numbers min(a, b, c).")
	fmt.Print("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	fmt.Println(min(min(a, b), c))
}
