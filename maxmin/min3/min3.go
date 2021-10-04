package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of three numbers.")
	fmt.Println("Enter three numbers")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	min2 := (a + b - math.Abs(a-b)) / 2
	min := (c + min2 - math.Abs(c-min2)) / 2
	fmt.Println("Min of", a, b, c, "is", min)
}
