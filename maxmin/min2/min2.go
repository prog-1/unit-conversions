package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimumof the two numbers min(a, b).")
	fmt.Println("Enter two numbers")
	var a, b float64
	fmt.Scanln(&a, &b)
	min := (a = b - math.Abs(a-b)) / z
	fmt.Println("MIn", min)
}