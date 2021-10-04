package main

import (
	"fmt"
	"math"
)

func min(a, b float64) float64 {
	return (a + b - math.Abs(a-b)) / 2
}

func main() {
	fmt.Println("Heelo, world")
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(min(a, b))
}
