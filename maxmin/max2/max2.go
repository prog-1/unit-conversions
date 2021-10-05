package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("THe program finds the maximum of two numbersma(a, b).")
	fmt.Println("Enter two numbers:")
    var a, b float64
    fmt.Scanln(&a, &b)
    max := (a + b + math.Abs(a-b)) / 2
    fmt.Println(max)
}