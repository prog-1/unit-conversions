package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program calculates the minimum of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Println("min:", (a+b-math.Abs(a-b))/2)
}
