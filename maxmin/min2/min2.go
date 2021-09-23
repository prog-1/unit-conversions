package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("This program show minimum of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scan(&a, &b)
	min := (a + b - math.Abs(a-b)) / 2
	fmt.Println("The mminimum of two numbers is:", min)
}
