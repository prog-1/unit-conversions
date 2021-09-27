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
	fmt.Println("Enter km")
	var k float64
	fmt.Scanln(&k)
	fmt.Println("km in m:", k*1000/3600)
}
