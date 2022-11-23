/* channel behaivours */

/*
modify the program in such a way that the "RECEIVE" operation & print is performed in a goroutine
*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 100
	data := <-ch
	fmt.Println(data)
}
