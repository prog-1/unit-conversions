package main

import "fmt"

func main() {
	fmt.Println("The program converts km/h to m/s.")
	fmt.Println("Enter the speed in km/h:")
	var speedkm float64
	fmt.Scanln(&speedkm)
	speedm := speedkm / 3.6
	fmt.Println("The speed in m/s:", speedm)
}
