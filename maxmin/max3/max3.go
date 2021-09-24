package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a > b && a > c {
		fmt.Println("Max:", a)
	} else if b > c {
		fmt.Println("Max:", b)
	} else {
		fmt.Println("Max:", c)
	}
}
