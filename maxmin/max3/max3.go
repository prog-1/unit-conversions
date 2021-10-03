package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers max(a, b, c).")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	maxab := (a + b + math.Abs(a-b)) / 2
	maxabc := (maxab + c + math.Abs(maxab-c)) / 2
	fmt.Println("max:", maxabc)
}
