package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("program that calculates the minimum of two numbers")
	fmt.Println("write two numbers:")
	var a, b float64
	fmt.Scan(&a, &b)
	minimum := (5 + 7 - math.Abs(a-b)) / 2
	fmt.Println("minimum of numbers:", minimum)
}
