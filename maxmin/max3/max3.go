package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers max(a, b, c).")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	max2 := (a + b + math.Abs(a-b)) / 2
	max := (c + max2 + math.Abs(c-max2)) / 2
	fmt.Println("Maximum", max)
}