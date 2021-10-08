package main

import "fmt"

func main() {
	fmt.Println("The program converts km/h to m/s.")
	var kmh float64
	fmt.Println("Enter the speed in km/h:")
	fmt.Scan(&kmh)
	fmt.Println("The speed in m/s", kmh/3.6)
}
