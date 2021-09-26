package main

/* (maxmin/max2) a program that calculates the maximum of three numbers.*/

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers.")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	//max := math.Max(math.Max(a, b), c)
	max1 := (a + b + math.Abs(a-b)) / 2
	max2 := (b + c + math.Abs(b-c)) / 2
	max := (max1 + max2 + math.Abs(max1-max2)) / 2
	fmt.Printf("max: %v", max)
}
