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
	// ++ Your code here! ++
	var a float64
	fmt.Print("Enter the speed in km/h: ")
	fmt.Scan(&a)
	fmt.Println("The speed in m/s:", (a / 3.6))
}
