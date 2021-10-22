package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	var a ,b float64
	fmt.Println("Enter the resistance of two resistors: ")
	fmt.Scan(&a, &b)
	//????????
	//the resistors R1 and R2 are connected in series, i.e. R = R1 + R2;
	//the resistors are connected in parallel, i.e. 1 / R = 1 / R1 + 1 / R2.
	r := a + b
	r1 := 1 / a + 1 / b
	r2 := 1 / r1
	fmt.Println(r ,r1 ,r2 )
}