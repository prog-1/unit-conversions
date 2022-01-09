package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("quadratic equation")
	var a, b, c float64
	fmt.Println("Enter A B C:")
	fmt.Scan(&a, &b, &c)
	D := b*b - 4*a*c
	fmt.Println("x1=", (-b+math.Sqrt(D))/(2*a))
	fmt.Println("x2=", (-b-math.Sqrt(D))/(2*a))

}
