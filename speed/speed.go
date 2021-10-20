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
	var sk float64
	fmt.Println("Enter speed km/h: ")
	fmt.Scan(&sk)
    spe = sk * 1000
	fmt.Println("Speed is ", spe)
}
