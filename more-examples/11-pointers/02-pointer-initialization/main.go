package main

import "fmt"

func main() {
	var a = []int {1,2,4,5,6}
	var p = &a


	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p = ", p)
}
