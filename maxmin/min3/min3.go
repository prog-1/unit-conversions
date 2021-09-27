package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the mnimum of three numbers")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	m := (a + b - math.Abs(a-b)) / 2
	min := (m + c - math.Abs(m-c)) / 2
	fmt.Println("Min:", min)
}
