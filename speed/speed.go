package main

import "fmt"

/*

Sample input/output:

The program converts km/h to m/s.
Enter the speed in km/h:
100
The speed in m/s: 27.77777777777778

*/
func main() {
	fmt.Println("The program converts km/h to m/s.")
	var kmph float64
	fmt.Println("Enter the speed in km/h:")
	fmt.Scanln(&kmph)
	var mps float64
	mps = kmph / 3.6
	fmt.Println("Speed in m/s:", mps)
}