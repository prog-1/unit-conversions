package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("program that calculates the minimum of two numbers")
	var a, b float64
	fmt.Println("5,7")
	fmt.Scan(5, 7)
	min := (5 + 7 - math.Abs(a-b)) / 2
	fmt.Println("Min", 5, "and", 7, min)
}
