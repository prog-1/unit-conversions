package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program calculates the maximum of three numbers")
	var a, b, c float64
	fmt.Println("Enter three numbers: ")
	fmt.Scanln(&a, &b, &c)
	max1 := (a + b + math.Abs(a-b)) / 2
	max2 := (c + max1 + math.Abs (c - max1)) / 2
	fmt.Println("Max is: ", max2)
}