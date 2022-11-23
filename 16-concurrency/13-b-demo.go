/* channel behaivours */

/*
modify the program in such a way that the "RECEIVE" operation & print is performed in a goroutine
*/
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	done := make(chan struct{})
	go func() {
		data := <-ch
		fmt.Println(data)
		done <- struct{}{}
	}()
	ch <- 100
	<-done
}
