package main

import "fmt"

func main() {
	// Initializing a map
	var tinderMatch = make(map[string]string)

	// Adding keys to a map
	tinderMatch["RCB"] = "kohli" // Assigns the value "Angelina" to the key "Rajeev"
	tinderMatch["CSK"] = "Dhoni"
	tinderMatch["Delhi"] = "Panth"

	fmt.Println(tinderMatch)

	/*
	  Adding a key that already exists will simply override
	  the existing key with the new value
	*/
	tinderMatch["CSK"] = "Jadeja"
	fmt.Println(tinderMatch)
}
