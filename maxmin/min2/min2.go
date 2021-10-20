package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minium of two numbers max(a, b).")
	var a, b float64
	fmt.Println("Enter numbers: ")
	fmt.Scan(&a, &b)
	min := (a + b - math.Abs(a-b)) / 2
	fmt.Println("Min is ", min)
}