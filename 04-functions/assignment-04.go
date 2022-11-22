package main

import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	//counter = 10000
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func getIncrement() func() int {
	var counter int = 0
	return func() int {
		counter++
		return counter
	}
}
