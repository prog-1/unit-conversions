package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("This program outputs min and max numbers")
	var a, b float64
	fmt.Println("enter 2 numbers")
	fmt.Scan(&a, &b)
	max := (a + b + math.Abs(a-b)) / 2
	fmt.Println("Max of", a, "and", b, "is", max)

}
