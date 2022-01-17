package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds the maximum of three numbers max(a, b, c).")
	fmt.Println("Enter three numbers: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	fmt.Println("Max is", (c+((a+b+math.Abs(a-b))/2)+math.Abs(c-(a+b+math.Abs(a-b))/2))/2)

}
