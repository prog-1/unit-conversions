package main

import (
	"fmt"
	"math"
) 

func main() {
	fmt.Println("The program finds the maximum of two numbers max(a, b).")
	var a ,b ,max float64
	fmt.Println("Enter numbers: ")
	fmt.Scan(&a, &b)
	//max(a, b) = (a + b + |a - b|)/2
	max := (a + b + math.Abs(a-b))/2
	fmt.Println("max: ", max)

}
