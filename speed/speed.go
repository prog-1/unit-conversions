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
	var v float64
	fmt.Scanln(&v)
	fmt.Printf("The speed in m/s: %.2f", v*10/36)
	// ++ Your code here! ++
}
