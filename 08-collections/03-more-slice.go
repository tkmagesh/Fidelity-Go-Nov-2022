package main

import "fmt"

func main() {
	// var list []int

	/* pre-allocating memory */
	var list []int = make([]int, 0, 10)

	/* intializing(5) and pre-allocating(10) */
	//var list []int = make([]int, 5, 10)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
	list = append(list, 10)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
	list = append(list, 20)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
	list = append(list, 30)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
	list = append(list, 40)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
	list = append(list, 50)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
	list = append(list, 60)
	fmt.Printf("len(list) = %d, cap(list) = %d, list = %v\n", len(list), cap(list), list)
}
