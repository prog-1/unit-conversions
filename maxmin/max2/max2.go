package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("This program show the max of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scan(&a, &b)
	max := (a + b + math.Abs(a-b)) / 2
	fmt.Println("The max of two numbers is:", max)
}
