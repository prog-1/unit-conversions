package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("program that calculates the maximum of three numbers")
	fmt.Println("Write three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	x := (b + a + math.Abs(a-b)) / 2
	maximum := (c + x + math.Abs(c-x)) / 2
	fmt.Println("Max", maximum)
}
