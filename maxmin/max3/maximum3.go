package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers max (a, b, c).")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	maxof2 := (a + b + math.Abs(a-b)) / 2
	fmt.Println("Max is", (c+maxof2+math.Abs(c-maxof2))/2)

}
