package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("program that calculates the minimum of three numbers ")
	fmt.Println("Write three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	x := (a + b - math.Abs(a-b)) / 2
	minimum := (x + c - math.Abs(x-c)) / 2
	fmt.Println("Minimum:", minimum)

}
