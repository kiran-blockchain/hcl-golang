package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)
	firstName ="Kiran"
	lastName="Paturi"
	isConfirmed =true

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)
}
