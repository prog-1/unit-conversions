package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of two numbers max(a, b).")
	var a, b float64
	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)
	fmt.Println("Max of", a, "and", b, "is", (a+b+math.Abs(a-b))/2)
}
