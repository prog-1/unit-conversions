package main

/*  a program that calculates the minimum of three numbers.*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of three numbers.")
	fmt.Println("Enter three numbers:")
	var a float64
	var b float64
	var c float64
	fmt.Scan(&a, &b, &c)
	var min1 = math.Min(a, b)
	var min2 = math.Min(b, c)
	var min = math.Min(min1, min2)
	fmt.Printf("min: %v", min)
}
