package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go genFibanocci(ch)
	for no := range ch {
		fmt.Println(no)
	}
}

func genFibanocci(ch chan int) {
	timeoutCh := timeout(10 * time.Second)
	x, y := 0, 1
LOOP:
	for {
		select {
		case <-timeoutCh:
			break LOOP
		default:
			time.Sleep(500 * time.Millisecond)
			ch <- x
			x, y = y, x+y
		}
	}
	close(ch)
}

func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
