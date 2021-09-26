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
	fmt.Println("Enter speed km/h")
	var v float64
	fmt.Scanln(&v)
	ms := v / 3.6
	fmt.Println(v, "km/h =", ms, "m/s")
	// ++ Your code here! ++
}
