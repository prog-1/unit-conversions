package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	var a, b, c float64
	fmt.Print("Enter the coefficients A, B and C: ")
	fmt.Scan(&a, &b, &c)
	D := b*b - 4*a*c
	fmt.Println("x1=", (-b+math.Sqrt(D))/(2*a))
	fmt.Println("x2=", (-b-math.Sqrt(D))/(2*a))

}
