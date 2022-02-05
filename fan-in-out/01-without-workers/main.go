package main

import (
	"fmt"
	"time"
)

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func PrepareItems(done <-chan bool) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		Item{0, "Shirt", 1 * time.Second},
		Item{1, "Legos", 1 * time.Second},
		Item{2, "TV", 5 * time.Second},
		Item{3, "Bananas", 2 * time.Second},
		Item{4, "Hat", 1 * time.Second},
		Item{5, "Phone", 2 * time.Second},
		Item{6, "Plates", 3 * time.Second},
		Item{7, "Computer", 5 * time.Second},
		Item{8, "Pint Glass", 3 * time.Second},
		Item{9, "Watch", 2 * time.Second},
	}
	go func() {
		for _, item := range itemsToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()
	return items
}

func PackItems(done <-chan bool, items <-chan Item) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
			}
		}
		close(packages)
	}()
	return packages
}

func main() {
	done := make(chan bool)
	defer close(done)

	start := time.Now()

	packages := PackItems(done, PrepareItems(done))

	numPackages := 0
	for p := range packages {
		numPackages++
		fmt.Printf("Shipping package no. %d\n", p)
	}

	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
