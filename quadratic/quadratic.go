package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	var a, b, c float64
	fmt.Println("Enter the coefficients A, B and C: ")
	fmt.Scan(&a, &b, &c)
	//ax^2 + bx + c
	d := b * b - 4 * a * c
	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)
	fmt.Println("x1 = ",x1 ,"x2 = ",x2)
}