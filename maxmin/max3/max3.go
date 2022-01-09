package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(" the maximum of three numbers")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	maxab := (a + b + math.Abs(a-b)) / 2
	maxabc := (maxab + c + math.Abs(maxab-c)) / 2
	fmt.Println("max:", maxabc)
}
