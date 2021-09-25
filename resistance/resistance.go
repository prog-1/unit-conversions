package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var R1 float64
	var R2 float64
	fmt.Scan(&R1, &R2)
	var RConnected = R1 + R2
	var Rparalel = (R1 * R2) / (R1 + R2)
	fmt.Printf("Resistance when connected in series is: %v ; resistance when connected in parallel is %v", RConnected, Rparalel)
}
