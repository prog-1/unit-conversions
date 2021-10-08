package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program prints the perimeter and the square of a circle given the radius.")
	var r float64
	fmt.Print("Enter the radius:")
	fmt.Scan(&r)
	fmt.Println("Perimeter:", r*2*math.Pi)
	fmt.Println("Square:", r*r*math.Pi)
}
