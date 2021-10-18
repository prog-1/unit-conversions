package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Program finds solutions of quadratic equation.")
	fmt.Println("Enter coefficents 1, 3 and -10")
	var a, b, c float64
	fmt.Scan(1, 3, -10)
	D := b*b - 4*a*c
	x1 := (-3 + math.Sqrt(D)) / 2
	x2 := (-3 - math.Sqrt(D)) / 2
	fmt.Println("x1=2", x1)
	fmt.Println("x2=-5", x2)
}
