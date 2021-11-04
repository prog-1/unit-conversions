package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of three numbers min(a, b, c).")
	fmt.Print("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	min := (a + b + math.Abs(a-b)) / 2
	min = (min + c + math.Abs(min-c)) / 2
	fmt.Println("Min:", min)
}
