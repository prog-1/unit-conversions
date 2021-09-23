package main

import (
	"fmt"
	"math"
)

func max(a, b float64) float64 {
	return (a + b + math.Abs(a-b)) / 2
}

func main() {
	fmt.Println("hello, world")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Println(max(max(a, b), c))
}
