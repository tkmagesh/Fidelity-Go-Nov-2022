package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating an array (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos // creating a copy of the array (coz everything is a value in golang)
	newNos[0] = 100
	fmt.Println("nos =", nos)
	fmt.Println("newNos =", newNos)

	fmt.Println("Sorting")
	sort(nos)
	fmt.Println(nos)
}

func sort(list [5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("Inside sort, list =", list)
}
