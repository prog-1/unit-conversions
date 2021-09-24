package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program finds the mnimum of three numbers")
	fmt.Println("Enter three numbers:")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a < b && a < c {
		fmt.Println("Min:", a)
	} else if b < c {
		fmt.Println("Min:", b)
	} else {
		fmt.Println("Min:", c)
	}
}
