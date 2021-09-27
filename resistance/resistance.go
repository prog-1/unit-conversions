package main

import "fmt"

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var r1, r2 float64
	fmt.Scan(&r1, &r2)
	rs := r1 + r2
	rp := 1/r1 + 1/r2
	fmt.Println("Resistance when connected in series is:", rs)
	fmt.Println("Resistance when connected in parallel is:", 1/rp)
}
