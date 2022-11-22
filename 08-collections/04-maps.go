package main

import (
	"fmt"
)

func main() {
	// var productRanks map[string]int
	// var productRanks map[string]int = map[string]int{"Pen": 3, "Pencil": 2, "Marker": 5}
	// productRanks := map[string]int{"Pen": 3, "Pencil": 2, "Marker": 5}
	/*
		productRanks := map[string]int{
			"Pen":    3,
			"Pencil": 2,
			"Marker": 5,
		}
	*/
	productRanks := make(map[string]int)
	productRanks["Pen"] = 3
	productRanks["Pencil"] = 2
	productRanks["Marker"] = 5

	fmt.Println(productRanks)

	fmt.Println("Adding a new key/value pair")
	productRanks["Scribble-Pad"] = 4
	fmt.Println("Count of product ranks =", len(productRanks))
	fmt.Println(productRanks)

	fmt.Println("Iterating a map (using range)")
	for key, val := range productRanks {
		fmt.Printf("ProductRanks[%q] = %d\n", key, val)
	}

	keyToCheck := "Pen"
	if rank, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key %q does not exist\n", keyToCheck)
	}

	keyToDelete := "Notepad"
	delete(productRanks, keyToDelete)
	fmt.Printf("After deleting %q, productRanks = %v\n", keyToDelete, productRanks)

}
