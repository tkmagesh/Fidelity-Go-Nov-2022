package main

import "fmt"

func main() {
	var x int = 10
	/*
		var f float32 = 12.99
		fmt.Println(&f)
	*/
	var xPtr *int = &x
	fmt.Println(xPtr)

	//dereferencing - accessing the value using the pointer
	fmt.Println(*xPtr)
	fmt.Println(x == *(&x))

	var no int = 100
	fmt.Println("Before incrementing, no =", no)
	increment(&no)
	fmt.Println("After incrementing, no =", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swapping, n1 = %d and n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("After swapping, n1 = %d and n2 = %d\n", n1, n2)
}

func increment(n *int) {
	*n++
}

func swap(n1, n2 *int) {
	*n1, *n2 = *n2, *n1
}
