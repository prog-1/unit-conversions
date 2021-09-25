package main

/*  a program that calculates the minimum of two numbers. */

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of two numbers.")
	fmt.Println("Enter two numbers:")
	var a float64
	var b float64
	fmt.Scan(&a, &b)
	var min = math.Min(a, b)
	fmt.Printf("min: %v", min)
}
