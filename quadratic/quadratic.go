package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("program that prints the equation solutions:")
	fmt.Println("Enter a, b, c:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	D := b*b - (1 * 3 * -10)
	x1 := (-b + math.Sqrt(D)) / 2
	x2 := (-b - math.Sqrt(D)) / 2
	fmt.Println("x1 =", x1)
	fmt.Println("x2 =", x2)
}
