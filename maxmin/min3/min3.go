package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("the minimum of three numbers")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	minab := (a + b - math.Abs(a-b)) / 2
	minabc := (minab + c - math.Abs(minab-c)) / 2
	fmt.Println("min:", minabc)
}
