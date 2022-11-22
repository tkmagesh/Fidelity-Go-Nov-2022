package main

import "fmt"

func main() {
	// var nos []int
	// nos[0] = 10
	// nos = append(nos, 10)
	// nos = append(nos, 20, 30)

	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	nos = append(nos, 10)
	nos = append(nos, 20, 30)
	hundreds := []int{100, 200, 300, 400}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	fmt.Println("Iterating a slice (using indexer)")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating a slice (using range)")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	newNos := nos
	newNos[0] = 1000
	fmt.Println("nos =", nos)
	fmt.Println("newNos =", newNos)

	fmt.Println("After sorting")
	sort(nos)
	fmt.Println(nos)

	fmt.Println("slicing")
	fmt.Println("nos[2:6] =", nos[2:6]) // from index 2 to index 5 (6-1)
	fmt.Println("nos[2:] =", nos[2:])   // from index 2 to the end of the slice
	fmt.Println("nos[:6] =", nos[:6])   // from index 0 to 5 (6-1)

	//copying a slice
	var copyOfNos []int = make([]int, len(nos))
	fmt.Println("Before copying, copyOfNos =", copyOfNos)
	copy(copyOfNos, nos)
	fmt.Println("After copying, copyOfNos =", copyOfNos)
	nos[0] = 9000
	fmt.Println("nos =", nos)
	fmt.Println("copyOfNos =", copyOfNos)
}

func sort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
