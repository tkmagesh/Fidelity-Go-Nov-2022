package main

import (
	"fmt"
)

//Share memory BY communicating

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println("result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result //send the result to the channel\
}
