package main

import "fmt"

/*
Sample input/output:
The program converts km/h to m/s.
Enter the speed in km/h:
The speed in m/s: 27.77777777777778
*/
func main() {
	fmt.Println("The program converts km/h to m/s.")
	fmt.Println("Enter km/h:")
	var a float64
	fmt.Scanln(&a)
	fmt.Println("result:", a*1000/3600)
}
