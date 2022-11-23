package main

import (
	"fmt"
)

//Share memory BY communicating
//consumer
func main() {
	ch := add(100, 200)
	result := <-ch
	fmt.Println("result :", result)
}

//producer
func add(x, y int) <-chan int /* receive-only channel */ {
	ch := make(chan int)
	go func() {
		result := x + y
		ch <- result //send the result to the channel\
	}()
	return ch
}
