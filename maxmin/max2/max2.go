package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("the maximum of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scanln(&a, &b)
	max := (a + b + math.Abs(a-b)) / 2
	fmt.Println("max:", max)
}
