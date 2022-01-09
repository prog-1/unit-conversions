package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("the minimum of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scanln(&a, &b)
	min := (a + b - math.Abs(a-b)) / 2
	fmt.Println("min:", min)
}
