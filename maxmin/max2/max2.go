package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello, world")
	var a, b float64
	fmt.Scan(&a, &b)
	max := (a + b + math.Abs(a-b)) / 2
	fmt.Println(max)
}
