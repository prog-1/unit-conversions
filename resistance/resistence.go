package main

import (
	"fmt"
)

func main() {
	fmt.Println("the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println("Resistance when connected in series is", a+b)
	res := 1/a + 1/b
	res = 1 / res
	fmt.Println("Resistance when connected in parallel is", res)

}
