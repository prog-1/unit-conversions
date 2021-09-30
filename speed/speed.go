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
	fmt.Println("Enter the speed in km/h:")
	var kmh float64
	fmt.Scan(&kmh)
	var ms float64 = kmh / 3.6
	fmt.Printf("The speed in m/s: %v", ms)
}
