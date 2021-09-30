package main

/*  a program that calculates the minimum of three numbers.*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of three numbers.")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	//var min = math.Min(math.Min(a, b), c)
	min1 := (a + b - math.Abs(a-b)) / 2
	min := (min1 + c - math.Abs(min1-c)) / 2
	fmt.Printf("min: %v", min)
}
