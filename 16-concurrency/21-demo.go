package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	ch := genFibanocci(done)

	fmt.Println("Hit ENTER to stop....")
	go func() {
		fmt.Scanln()
		done <- true
	}()

	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Thats all folks!!")
}

func genFibanocci(done chan bool) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- x
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}

/* func timeout(d time.Duration) <-chan time.Time {
	timeoutCh := make(chan time.Time)
	go func() {
		time.Sleep(d)
		timeoutCh <- time.Now()
	}()
	return timeoutCh
}
*/
