package main

import (
	"fmt"
	"math/rand"
	"time"
)

//consumer
func main() {
	ch := make(chan int)
	go fn(ch)

	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
			continue
		}
		break
	}
}

//producer
func fn(ch chan int) {
	rand.Seed(7)
	count := rand.Intn(20)
	for i := 1; i <= count; i++ {
		ch <- i * 10
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}
