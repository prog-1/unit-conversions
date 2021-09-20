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
	fmt.Println("Enter the speed in km/h: ")
	var x float64
	fmt.Scan(&x)
	x = x / 3.6
	fmt.Println("The speed in m/s:", x)

}
