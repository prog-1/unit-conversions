package main

import "fmt"

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var R1, R2 float64
	fmt.Scanln(&R1, &R2)
	R := R1 + R2
	RR := 1/R1 + 1/R2
	RRR := 1 / RR
	fmt.Print("Resistance when connected in series is ", R)
	fmt.Println("; resistance when connected in parallel is", RRR)
}
