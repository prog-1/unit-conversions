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
	min2 := (b + c - math.Abs(b-c)) / 2
	min := (min1 + min2 - math.Abs(min1-min2)) / 2
	fmt.Printf("min: %v", min)
}
