package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program calculates the maximum of two numbers")
	fmt.Println("Enter two numbers:")
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Println("max:", (a+b+math.Abs(a-b))/2)
}
