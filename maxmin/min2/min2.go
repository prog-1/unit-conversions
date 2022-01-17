package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of two numbers min(a, b).")
	fmt.Println("Enter two numbers: ")
	var a, b float64
	fmt.Scanln(&a, &b)
	fmt.Println("Min is", (a+b-math.Abs(a-b))/2)
}
