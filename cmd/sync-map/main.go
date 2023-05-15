package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Add items to the map
	m.Store("apple", 1)
	m.Store("banana", 2)
	m.Store("orange", 3)

	// Retrieve an item from the map
	value, ok := m.Load("apple")
	if ok {
		fmt.Println("Value found:", value)
	}

	// Iterate over the map
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, "->", value)
		return true // Continue iterating
	})

	// Delete an item from the map
	m.Delete("banana")
}
