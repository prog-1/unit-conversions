package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var R1, R2 float64
	fmt.Scanln(&R1, &R2)
	R := R1 + R2
	Rx := 1 / (R1 + R2)
	fmt.Println("Resistance connected in series:", R)
	fmt.Println("Resistance connected in Paralel:", Rx)

}
