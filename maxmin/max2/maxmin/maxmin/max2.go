package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Program calculates the maximum of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scan(5, 7)
	maximum := (5 + 7 + math.Abs(a-b)) / 2
	fmt.Println("maximum of numbers is:", maximum)
}
