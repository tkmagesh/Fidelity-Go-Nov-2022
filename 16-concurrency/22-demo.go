/* channel behaivours (buffered) */

package main

import "fmt"

func main() {
	ch := make(chan int, 1) //buffer size 1
	ch <- 100
	data := <-ch
	fmt.Println(data)
}
