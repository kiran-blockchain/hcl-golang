// Golang program to pass a pointer to an
// array as an argument to the function
package main

import "fmt"

// taking a function
func updatearray(funarr *[5]int) {

	// updating the array value
	// at specified index
	(*funarr)[4] = 750
	
	// you can also write
	// the above line of code
	// funarr[4] = 750
}

// Main Function
func main() {

	// Taking an pointer to an array
	arr := [5]int{78, 89, 45, 56, 14}

	// passing pointer to an array
	// to function updatearray
	updatearray(&arr)

	// array after updating
	fmt.Println(arr)
}
