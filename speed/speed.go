package main

import "fmt"

/*

Sample input/output:

The program converts km/h to m/s.
Enter the speed in km/h: 100
The speed in m/s: 27.78

*/
func main() {
	fmt.Println("The program converts km/h to m/s.")
	fmt.Print("Enter the speed in km/h: ")
	var kmh float64
	fmt.Scan(&kmh)
	fmt.Printf("The speed in m/s: %.2f", kmh/60/60*1000)
}
