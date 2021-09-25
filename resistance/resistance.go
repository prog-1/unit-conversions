package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Print("Enter the resistance of two resistors:")
	var R1, R2 float64
	fmt.Scan(&R1, &R2)
	var R, Rp float64
	R = R1 + R2
	Rp = (R1 * R2) / (R1 + R2)
	fmt.Println("Resistance when connected in series is:", R)
	fmt.Print("Resistance when connected in parallel is:", Rp)
}
