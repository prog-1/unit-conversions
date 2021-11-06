package main

import "fmt"

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Print("Enter the resistance of two resistors:")
	var R1, R2 float64
	fmt.Scan(&R1, &R2)
	S := R1 + R2
	P := 1/R1 + 1/R2
	P = 1 / P
	fmt.Println("Resistance when connected in series is", S, "; resistance when connected in parallel is", P)
}
