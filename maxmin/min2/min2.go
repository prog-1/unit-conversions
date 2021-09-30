package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of two numbers min(a, b).")
	var a, b float64
	fmt.Println("Enter two numbers")
	fmt.Scanln(&a, &b)
	fmt.Println("Min of", a, "and", b, "is", (a+b-math.Abs(a-b))/2)
}
