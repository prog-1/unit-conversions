package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of two numbers max(a, b).")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scanln(&a, &b)
	max := (a + b + math.Abs(a-b)) / 2
	fmt.Println(max)
}