package main

import "fmt"

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("enter R1 and R2 resistance")
	var r1, r2 float64
	fmt.Scanln(&r1, &r2)
	R := r1 + r2
	rr := 1 / (1/r1 + 1/r2)
	fmt.Println("Resistance when connected in series is", R)
	fmt.Println("Resistance when connected in parallel is", rr)

}
