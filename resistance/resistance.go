package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	var R1, R2 float64
	fmt.Println("Enter the resistance of two resistors:")
	fmt.Scan(&R1, &R2)
	fmt.Println("Resistance when connected in series is", R1+R2)
	pinv := 1/R1 + 1/R2
	pinv = 1 / pinv
	fmt.Println("Resistance when connected in parallel is", pinv)

}
