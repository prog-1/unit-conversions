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
	var speedkmh float64
	fmt.Print("Enter the speed in km/h:")
	fmt.Scan(&speedkmh)
	var speedms float64
	speedms = speedkmh / 3.6
	fmt.Println("The speed in m/s:", speedms)
}
