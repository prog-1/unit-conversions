package main

import "fmt"

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var R1, R2 float64
	fmt.Scan(&R1, &R2)
	series := R1 + R2
	parallel := 1/R1 + 1/R2
	parallel = 1 / parallel
	fmt.Println("Resistance when connected in series is", series, "resistance when connected in parallel is", parallel)
}
