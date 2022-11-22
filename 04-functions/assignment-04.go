package main

import "fmt"

var counter int = 0

func main() {
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	counter = 10000
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func increment() int {
	counter++
	return counter
}
