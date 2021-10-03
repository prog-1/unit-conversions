package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	fmt.Println("Enter the radius:")
	var r float64
	fmt.Scanln(&r)
	p := 2 * math.Pi * r
	s := math.Pi * r * r
	fmt.Println("Perimeter:", p)
	fmt.Println("Square:", s)
}
