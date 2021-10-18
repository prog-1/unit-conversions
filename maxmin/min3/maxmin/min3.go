package main

import (
	"fmt"
)

func main() {
	fmt.Println("program that calculates the minimum of three numbers ")
	fmt.Println("5, 7, 10:")
	var a, b, c float64
	fmt.Scan(5, 7, 10)

	fmt.Println("min=", min(min(a, b), c))
}
