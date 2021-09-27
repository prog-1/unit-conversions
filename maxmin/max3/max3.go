package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program show max of 3 numbers")
	fmt.Print("Enter 3 numbers: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	m := (a + b + math.Abs(a-b)) / 2
	max := (c + m + math.Abs(c-m)) / 2
	fmt.Println("Max:", max)
}
