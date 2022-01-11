package main

import "fmt"

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	var a, b float64
	fmt.Print("Enter the resistance of two resistors: ")
	fmt.Scan(&a, &b)
	fmt.Println("Resistance when connected in series is ", (a + b), ";", "resistance when connected in parallel is", (1 / (1/a + 1/b)))
}
