package main

/* (maxmin/max2) a program that calculates the maximum of three numbers.*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers.")
	fmt.Println("Enter three numbers:")
	var a float64
	var b float64
	var c float64
	fmt.Scan(&a, &b, &c)
	var max1 = math.Max(a, b)
	var max2 = math.Max(b, c)
	var max = math.Max(max1, max2)
	fmt.Printf("max: %v", max)
}
