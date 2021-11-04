package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers max(a, b, c).")
	fmt.Print("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	max := (a + b + math.Abs(a-b)) / 2
	max = (max + c + math.Abs(max-c)) / 2
	fmt.Println("Max:", max)
}
