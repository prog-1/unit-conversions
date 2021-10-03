package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program calculates the minimum of three numbers")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	min2 := (a + b - math.Abs(a-b)) / 2
	min := (c + min2 - math.Abs(c-min2)) / 2
	fmt.Println("Minimum:", min)
}
