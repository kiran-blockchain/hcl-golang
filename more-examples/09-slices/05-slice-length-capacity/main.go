package main

import "fmt"

func main() {
	/*
		The length of the slice is the number of elements in the slice.
		The capacity is the number of elements in the underlying array starting from the first element in the slice.
	*/
	a := [7]int{10, 20, 30, 40, 50, 60,90}
	s := a[1:2]

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
