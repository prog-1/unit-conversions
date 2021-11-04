package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of two numbers max(a, b).")
	fmt.Print("Enter two numbers:")
	var a, b float64
	fmt.Scan(&a, &b)
	max := (a + b + math.Abs(a-b)) / 2
	fmt.Println("Max:", max)
}
