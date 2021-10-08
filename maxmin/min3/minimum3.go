package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the minimum of three numbers min(a, b, c).")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	minof2 := (a + b - math.Abs(a-b)) / 2
	fmt.Println("Min is", (c+minof2-math.Abs(c-minof2))/2)

}
